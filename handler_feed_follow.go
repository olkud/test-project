package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/olkud/test-project/internal/database"
	"net/http"
	"time"
)

func (api *apiConfig) handlerFollowFeed(w http.ResponseWriter, r *http.Request, user User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	decodeError := decoder.Decode(&params)

	if decodeError != nil {
		respondWithJSON(w, 400, fmt.Sprintln("Error passing JSON: ", decodeError))
		return
	}

	feedFollow, err := api.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		responseWithError(w, 400, fmt.Sprintln("Could not create feed: ", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feedFollow))
}

func (api *apiConfig) handlerGetUserFeeds(w http.ResponseWriter, r *http.Request, user User) {
	feedFollows, err := api.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		responseWithError(w, 400, fmt.Sprintln("Could not get user feeds: ", err))
		return
	}

	respondWithJSON(w, 200, databaseFeedFollowsToFeedFollows(feedFollows))
}

func (api *apiConfig) handlerUnfollowFeed(w http.ResponseWriter, r *http.Request, user User) {
	feedFollowId := chi.URLParam(r, "feedFollowID")

	if feedFollowId == "" {
		respondWithJSON(w, 400, fmt.Sprintln("Invalid feedFollowID"))
		return
	}

	uuidId, err := uuid.Parse(feedFollowId)

	if err != nil {
		respondWithJSON(w, 400, fmt.Sprintln("feedFollowID is not UUID"))
	}

	err = api.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     uuidId,
		UserID: user.ID,
	})

	if err != nil {
		responseWithError(w, 400, fmt.Sprintln("Could not unfollow feed: ", err))
		return
	}

	respondWithJSON(w, 204, struct{}{})
}
