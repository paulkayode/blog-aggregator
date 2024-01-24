package main

import "net/http"

func readinessHandler(w http.ResponseWriter, r * http.Request){
	type response struct {
		Status string `json:"status"`
	}
	resp := response{
		Status : "ok",
	} 
	respondWithJson(w, 200, resp)
}