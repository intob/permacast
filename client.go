package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	arHost = "https://arweave.net"
)

type Query struct {
	Query string `json:"query"`
}

type Result struct {
	Data Data `json:"data"`
}

type Data struct {
	Transactions Transactions `json:"transactions"`
}

type Transactions struct {
	Edges []Edge `json:"edges"`
}

type Edge struct {
	Node Node `json:"node"`
}

type Node struct {
	Id   string `json:"id"`
	Tags []Tag  `json:"tags"`
}

type Tag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func makeQuery(txId string) (*Result, error) {
	q := &Query{fmt.Sprintf(`query {
			transactions (
				ids: ["%s"]
				tags: [
					{ name: "Protocol", values: ["permacast"] }
				]) {
				edges {
					node {
						id
						tags {
							name
							value
						}
					}
				}
			}
	}`, txId)}

	qb, err := json.Marshal(q)
	if err != nil {
		return nil, err
	}

	qbr := bytes.NewReader(qb)

	response, err := http.Post(fmt.Sprintf("%s/graphql", arHost), "application/json", qbr)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %s", err)
	}

	res := &Result{}
	err = json.Unmarshal(buf.Bytes(), res)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal query response: %s", err)
	}
	return res, nil
}
