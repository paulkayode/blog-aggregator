package main 

import (
	"net/http"
	"context"
	"log"
)
func (cfg * apiConfig) GetFeedsHandler(w http.ResponseWriter, r * http.Request){
	ctx := context.Background()
	feeds, err := cfg.DB.GetFeeds(ctx)

	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	resp := []Feed{};
	for _, feed := range feeds{
		resp = append(resp, GetFeedFromDatabaseFeed(feed))
	}
	respondWithJson(w, http.StatusOK, resp)
}