package main

import (
	"fmt"
	"github.com/olkud/test-project/internal/auth"
	"net/http"
)

type authedHandler func(w http.ResponseWriter, r *http.Request, user User)

func (api *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)

		if err != nil {
			responseWithError(w, 403, fmt.Sprintln("Could not auth user"))
			return
		}

		user, err := api.DB.GetUserByAPIKey(r.Context(), apiKey)

		if err != nil {
			responseWithError(w, 403, fmt.Sprintln("Could not get user"))
			return
		}

		handler(w, r, databaseUserToUser(user))
	}
}
