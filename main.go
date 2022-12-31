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
		log.Printf("GET %s", r.URL.Path)
		http.Error(w, "method not supported", http.StatusBadRequest)
		return
	}
	handleGetData(w, r)
}
