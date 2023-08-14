package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/olkud/test-project/internal/database"
	"net/http"
	"time"
)

func (api *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)

	defer r.Body.Close()
	params := parameters{}

	decodeError := decoder.Decode(&params)

	if decodeError != nil {
		respondWithJSON(w, 400, fmt.Sprintln("Error passing JSON: ", decodeError))
		return
	}

	feed, err := api.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})

	if err != nil {
		responseWithError(w, 400, fmt.Sprintln("Could not create feed: ", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}

func (api *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := api.DB.GetFeeds(r.Context())

	if err != nil {
		responseWithError(w, 400, fmt.Sprintln("Could not get feeds: ", err))
		return
	}

	respondWithJSON(w, 200, databaseFeedsToFeeds(feeds))
}
