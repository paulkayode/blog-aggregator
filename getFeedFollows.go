package main

import (
	"context"
	"log"
	"net/http"
	"github.com/segunkayode1/blog-aggregator/internal/database"
)

func (cfg * apiConfig) GetFeedFollowsHandler(w http.ResponseWriter, r * http.Request,user database.User){
	ctx := context.Background()
	feedfollows, err := cfg.DB.GetAllFeedFollowsForUser(ctx, user.ID)

	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	resp := []UsersFeed{};
	for _, userfeed := range feedfollows{
		resp = append(resp, GetUsersFeedsFromDatabaseUsersFeeds(userfeed))
	}
	respondWithJson(w, http.StatusOK, resp)
}