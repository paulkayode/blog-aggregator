package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"context"
	"github.com/google/uuid"
	"github.com/segunkayode1/blog-aggregator/internal/database"
)

  
func (cfg *apiConfig)PostUser(w http.ResponseWriter, r * http.Request){
	ctx := context.Background()
	type paramters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	req := paramters{}
	err := decoder.Decode(&req)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	Useruuid := uuid.New()

	now := time.Now()
    user, err := cfg.DB.CreateUser(ctx, database.CreateUserParams{
		ID: Useruuid,
		CreatedAt: now,
		UpdatedAt: now,
		Name: req.Name,
	})
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
    resp := GetUsserFromDatabaseUser(user)
	respondWithJson(w, http.StatusOK, resp)
}