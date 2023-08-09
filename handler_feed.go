package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/olkud/test-project/internal/database"
	"net/http"
	"time"
)

func (api *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	decodeError := decoder.Decode(&params)

	if decodeError != nil {
		respondWithJSON(w, 400, fmt.Sprintln("Error passing JSON: ", decodeError))
		return
	}

	user, err := api.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		responseWithError(w, 400, fmt.Sprintln("Could not create user: ", err))
		return
	}

	respondWithJSON(w, 201, databaseUserToUser(user))
}

func (api *apiConfig) handlerGetUserByAPIKey(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, databaseUserToUser(user))
}
