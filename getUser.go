package main

import (
	"net/http"
	"github.com/segunkayode1/blog-aggregator/internal/database"
)

  
func (cfg *apiConfig)GetUserHandler(w http.ResponseWriter, r * http.Request, user database.User){
    resp := GetUserFromDatabaseUser(user)
	respondWithJson(w, http.StatusOK, resp)
}