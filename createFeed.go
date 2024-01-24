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

  
func (cfg *apiConfig)PostFeedHandler(w http.ResponseWriter, r * http.Request, user database.User){
	ctx := context.Background()
	type paramters struct {
		Name string `json:"name"`
		Url string `json:"url"`
	}
	req := paramters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	now := time.Now()
	feed, err := cfg.DB.CreateFeed(ctx, database.CreateFeedParams{
		ID : uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name: req.Name,
		Url: req.Url,
		UserID: user.ID,
	})
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	feedfollow, err := cfg.DB.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID: user.ID,
		FeedID: feed.ID,
	})

	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp := struct {
		MyFeed Feed `json:"feed"`
		MyFeedFollow UsersFeed `json:"feed_follow"`
	}{
		MyFeed: GetFeedFromDatabaseFeed(feed),
		MyFeedFollow : GetUsersFeedsFromDatabaseUsersFeeds(feedfollow),
	}

	respondWithJson(w, http.StatusCreated, resp)

}