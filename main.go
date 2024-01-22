package main 

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"log"
)

func main(){
	//loading port from env
	godotenv.Load()
	port := os.Getenv("PORT")

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

	mainRouter.Mount("/v1", subRouterV1)
	
	server := http.Server{
		Addr: ":" + port,
		Handler : mainRouter,

	}
	log.Fatal(server.ListenAndServe())
}