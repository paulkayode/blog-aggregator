package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/segunkayode1/blog-aggregator/internal/database"
)

func main(){
	//loading port from env
	godotenv.Load()
	port := os.Getenv("PORT")
	dbUrl := os.Getenv("CONN")
    db,err := sql.Open("postgres",dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)
	cfg := &apiConfig{
		DB: dbQueries,
	}
	T := time.NewTicker(time.Minute);
	go func(t *time.Ticker){
		for now := range T.C{
			ctx := context.Background()
          fmt.Printf("Fetching feeds at %v\n", now)
		  feeds, err := cfg.DB.GetNextFeedToFetch(ctx, 10)
		  if err != nil {
			log.Println(err)
			continue;
		  }
		  wg := sync.WaitGroup{}
		  c := make(chan *returnVal, len(feeds))
		  for _, feed := range feeds {
			 wg.Add(1)
			 go cfg.GetRssData(feed.Url,feed.ID ,c, &wg)
			 cfg.DB.MarkFeedFetched(ctx, database.MarkFeedFetchedParams{
				LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
				UpdatedAt: time.Now(),
				ID:feed.ID,
			 })
		  }
		  wg.Wait()
		  close(c)

		  for feed := range c{
			  for _,post :=  range feed.val.Item {
					now = time.Now()
					_, err := cfg.DB.CreatePosts(ctx, database.CreatePostsParams{
						ID:uuid.New(),
						CreatedAt: now,
						UpdatedAt: now,
						Title: post.Title,
						Url: post.Link,
						Description: post.Description,
						PublishedAt: post.PubDate,
						FeedID: feed.id,
					})
					if err != nil {
						if err.Error() != "pq: duplicate key value violates unique constraint \"posts_url_key\""{
							log.Println(err)
						}
					}
			  }
		  }
		}
		
	}(T)
	mainRouter := chi.NewRouter()

	//cors middleWare 
	mainRouter.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	  }))

	//v1 endpoints
	subRouterV1 := chi.NewRouter()
	subRouterV1.Get("/readiness", readinessHandler)
	subRouterV1.Get("/err", errorHandler)
	subRouterV1.Post("/users", cfg.PostUserHandler)
	subRouterV1.Get("/users", cfg.MiddlewareAuth(cfg.GetUserHandler))
	subRouterV1.Post("/feeds", cfg.MiddlewareAuth(cfg.PostFeedHandler))
	subRouterV1.Get("/feeds", cfg.GetFeedsHandler)
	subRouterV1.Post("/feed_follows", cfg.MiddlewareAuth(cfg.PostFeedFollowHandler))
	subRouterV1.Delete("/feed_follows/{feedFollowID}", cfg.MiddlewareAuth(cfg.DeleteFeedFollowHandler))
	subRouterV1.Get("/feed_follows", cfg.MiddlewareAuth(cfg.GetFeedFollowsHandler))
	subRouterV1.Get("/posts", cfg.MiddlewareAuth(cfg.GetPosts))
	
	mainRouter.Mount("/v1", subRouterV1)
	

	server := http.Server{
		Addr: ":" + port,
		Handler : mainRouter,

	}
	log.Fatal(server.ListenAndServe())
}