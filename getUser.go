package main

import (
	"strings"
	"log"
	"net/http"
	"context"
)

  
func (cfg *apiConfig)GetUser(w http.ResponseWriter, r * http.Request){
	ctx := context.Background()

	apikey := r.Header.Get("Authorization")
	apikey = strings.TrimPrefix(apikey, "ApiKey ")
	user , err := cfg.DB.GetUserBYApiKey(ctx, apikey)
	if err != nil {
		log.Println(err)
		log.Println(apikey)
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

    resp := GetUsserFromDatabaseUser(user)
	respondWithJson(w, http.StatusOK, resp)
}