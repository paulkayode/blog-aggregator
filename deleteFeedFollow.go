package main

import (
	"context"
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/segunkayode1/blog-aggregator/internal/database"
)
func (cfg * apiConfig)DeleteFeedFollowHandler(w http.ResponseWriter, r * http.Request, user database.User){
	ctx := context.Background()
	feedFollowID := chi.URLParam(r, "feedFollowID")
	uuidFeedFollow , err := uuid.Parse(feedFollowID)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	 err2 := cfg.DB.DeleteFeedFollow(ctx, uuidFeedFollow )
	 if err2 != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	 }
	 resp := struct {
		Body string `json:"body"`
	}{
		Body: "Status Ok",
	}
	respondWithJson(w, http.StatusOK, resp)

}