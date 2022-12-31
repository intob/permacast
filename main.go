package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handleRoot)

	port := os.Getenv("PORT")
	if port == "" {
		port = "1992"
	}
	addr := ":" + port
	log.Printf("listening on %s\r\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method not supported", http.StatusBadRequest)
		return
	}
	// set default cache header
	w.Header().Set("Cache-Control", "public, max-age=3600")
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "landing/index.html")
		return
	}
	if r.URL.Path == "/styles.css" {
		http.ServeFile(w, r, "landing/styles.css")
		return
	}
	handleGetData(w, r)
}
