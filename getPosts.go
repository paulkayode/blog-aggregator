package main

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"github.com/segunkayode1/blog-aggregator/internal/database"
)

func (cfg *apiConfig)GetPosts(w http.ResponseWriter, r* http.Request, user database.User){
	ctx := context.Background()
	var limit int32 = 10
	if r.URL.Query().Has("limit") {
		stringlimit := r.URL.Query().Get("limit")
		val, err := strconv.Atoi(stringlimit)
		if err != nil {
			log.Println(err)
			respondWithError(w,http.StatusInternalServerError, err.Error())
			return
		}
		limit = int32(val)
	}
	posts, err := cfg.DB.GetPostsByUser(ctx, database.GetPostsByUserParams{
		UserID: user.ID,
		Limit: limit,
	})
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	resp := []Post{}
	for _,post := range posts {
		resp = append(resp, GetPostsFromDatabasePosts(post))
	}
	respondWithJson(w, http.StatusOK, resp)
}