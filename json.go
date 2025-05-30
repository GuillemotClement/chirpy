package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string, err error) {
	if err != nil {
		log.Println(err)
	}

	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}

	type ErrorResponse struct {
		Error string `json:"Error"`
	}

	respondWithJSON(w, code, ErrorResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(dat)
}

func respondEndpointForbidden(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-type", "application/json")
	errorMessage := map[string]string{
		"error": msg,
	}

	dat, err := json.Marshal(errorMessage)
	if err != nil {
		log.Printf("Error marshall JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(dat)
}
