package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"github.com/google/uuid"
	"github.com/segunkayode1/blog-aggregator/internal/database"
	"time"
)
func (cfg * apiConfig)PostFeedFollowHandler(w http.ResponseWriter, r * http.Request, user database.User){
	ctx := context.Background()
	type paramters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	req := paramters{}
	err := decoder.Decode(&req)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	now := time.Now()
	userfeed, err := cfg.DB.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID: user.ID,
		FeedID: req.FeedID,
	})
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, GetUsersFeedsFromDatabaseUsersFeeds(userfeed))

}