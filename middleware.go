package main 

import (
	"github.com/segunkayode1/blog-aggregator/internal/database"
	"net/http"
	"log"
	"strings"
	"context"
)
type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter,  r* http.Request){
		ctx := context.Background()
		apikey := r.Header.Get("Authorization")
	    apikey = strings.TrimPrefix(apikey, "ApiKey ")
	    user , err := cfg.DB.GetUserByApiKey(ctx, apikey)
	    if err != nil {
		log.Println(err)
		log.Println(apikey)
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	handler(w, r, user)
})
}
