package main

import (
	"fmt"
	"net/http"
)

func (api *apiConfig) handlerGetUserFollowedPosts(w http.ResponseWriter, r *http.Request, user User) {
	userPosts, err := api.DB.GetPostsForUser(r.Context(), user.ID)

	if err != nil {
		responseWithError(w, 400, fmt.Sprintln("Could not get user posts: ", err))
		return
	}

	respondWithJSON(w, 200, databaseGetPostsForUserRowToUserPosts(userPosts))
}
