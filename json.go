package main

import (
	"encoding/json"
	"net/http"
	"log"
)
func respondWithJson(w http.ResponseWriter, status int, val interface{}){
	data , err := json.Marshal(&val)
	if err != nil {
		log.Fatal(err)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
    w.Write(data)
}

func respondWithError(w http.ResponseWriter, status int, msg string){
	type errorResponse struct {
		Error string `json:"error"`
	}
	resp := errorResponse{
		Error: msg,
	}
	respondWithJson(w, status, resp)
}