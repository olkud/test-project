package main

import (
	"encoding/json"
	"fmt"
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
		responseWithError(w, 400, fmt.Sprintln("Could not create feed: ", err))
		return
	}

	respondWithJSON(w, 200, databaseFeedFollowsToFeedFollows(feedFollows))
}
