package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleGet(w, r)
		}
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "1992"
	}
	addr := ":" + port
	log.Printf("listening on %s\r\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	fileHash := strings.TrimPrefix(r.URL.Path, "/")
	if fileHash == "" {
		http.Error(w, "missing File-Hash in path", http.StatusBadRequest)
		return
	}
	qRes, err := makeQuery(fileHash)
	if err != nil {
		werr := fmt.Sprintf("failed to query Arweave: %s", err)
		http.Error(w, werr, http.StatusInternalServerError)
		log.Println(werr)
		return
	}
	if len(qRes.Data.Transactions.Edges) != 1 {
		http.Error(w, "nothing found for given File-Hash", http.StatusNotFound)
		return
	}
	node := qRes.Data.Transactions.Edges[0].Node
	contentType, err := getContentType(node.Tags)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	dataUrl := ""
	if contentType == "text/plain" {
		dataUrl = fmt.Sprintf("http://arweave.net/tx/%s/data", node.Id)
	} else {
		typeSegments := splitMimeType(contentType)
		subtype := ""
		if len(typeSegments) == 1 {
			subtype = typeSegments[0]
		} else if len(typeSegments) >= 2 {
			subtype = typeSegments[1]
		} else {
			http.Error(w, fmt.Sprintf("unable to handle Content-Type: %s", contentType), http.StatusInternalServerError)
			return
		}
		dataUrl = fmt.Sprintf("http://arweave.net/tx/%s/data.%s", node.Id, subtype)
	}
	resp, err := http.Get(dataUrl)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get transaction data: %s", err), http.StatusInternalServerError)
		return
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	w.Header().Add("Content-Type", contentType)
	w.Header().Add("Content-Length", strconv.Itoa(buf.Len()))
	w.Header().Add("Cache-Control", "public, max-age=604800, immutable")
	w.Write(buf.Bytes())
}

func getContentType(tags []Tag) (string, error) {
	for _, t := range tags {
		if t.Name == "Content-Type" {
			return t.Value, nil
		}
	}
	return "", errors.New("no Content-Type tag found")
}

func splitMimeType(mimeType string) []string {
	return strings.Split(mimeType, "/")
}
