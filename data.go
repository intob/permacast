package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func handleGetData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "public, max-age=604800, immutable")

	txId := strings.TrimPrefix(r.URL.Path, "/")
	if txId == "" {
		http.Error(w, "missing txId in path", http.StatusBadRequest)
		return
	}

	qRes, err := makeQuery(txId)
	if err != nil {
		werr := fmt.Sprintf("failed to query Arweave: %s", err)
		http.Error(w, werr, http.StatusInternalServerError)
		log.Println(werr)
		return
	}

	if len(qRes.Data.Transactions.Edges) == 0 {
		http.Error(w, "no transaction found for given txId", http.StatusNotFound)
		return
	}

	node := qRes.Data.Transactions.Edges[0].Node

	contentType, err := getContentType(node.Tags)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("unable to get Content-Type from tags: %s", err.Error())
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

	for _, t := range node.Tags {
		w.Header().Set(t.Name, t.Value)
	}
	w.Header().Set("Content-Length", strconv.Itoa(buf.Len()))
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
