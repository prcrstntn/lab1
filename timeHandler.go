package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func TimeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
    http.Error(w, "Method is not supported, use GET request", http.StatusNotFound)
    return
    } else if r.URL.Path != "/time" {
    http.Error(w, "404 page not found", http.StatusNotFound)
    return
  }
	
	current := time.Now().Format(time.RFC3339)
	response := map[string]string{"time": current}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		return
	}
}
