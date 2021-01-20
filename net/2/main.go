package main

import (
	"fmt"
	"net/http"

	stats_api "github.com/fukata/golang-stats-api-handler"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Hello, World!\n")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/stats", stats_api.Handler)
	http.ListenAndServe(":8080", nil)
}
