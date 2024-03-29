package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func customizations() {
	// Register any custom backend handlers needed
	http.HandleFunc("/api/time", timeHandler)

	// Register the paths you want to render the default index.html to enable SPA behavior
	allowedPathsForSPA = append(allowedPathsForSPA, "/time", "/counter")
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Time string `json:"time"`
	}{
		Time: time.Now().Format(time.RFC1123),
	})
}
