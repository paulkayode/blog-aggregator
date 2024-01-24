package main

import (
	"database/sql"
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
	/*
	"id": "3f8805e3-634c-49dd-a347-ab36479f3f83",
    "created_at": "2021-09-01T00:00:00Z",
    "updated_at": "2021-09-01T00:00:00Z",
    "name": "Lane"
	*/
	type response struct {
		Id string `json:"id"`
		CreatedAt string `json:"create_at"`
		UpdatedAt string `json:"updated_at"`
		Name string `json:"name"`
	}
	Useruuid := uuid.New()

	now := time.Now()
    cfg.DB.CreateUser(ctx, database.CreateUserParams{
		ID: Useruuid,
		CreatedAt: sql.NullTime{Time: now, Valid: true},
		UpdatedAt: sql.NullTime{Time: now, Valid: true},
		Name: sql.NullString{String:req.Name, Valid:true},
	})
    resp := response {
		Id : Useruuid.String(),
		CreatedAt: now.Format(time.RFC3339),
		UpdatedAt: now.Format(time.RFC3339),
		Name: req.Name,
	}
	respondWithJson(w, http.StatusOK, resp)
}