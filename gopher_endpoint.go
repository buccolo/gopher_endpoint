package main

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type WombatResponse struct {
	RequestId string   `json:"request_id"`
	Gophers   []Gopher `json:"gophers"`
}

type Gopher struct {
	Id int `json:"id"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	log.Printf("Starting gopher_endpoint...")

	http.HandleFunc("/", Root)

	var port = os.Getenv("PORT")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func extractRequestId(r *http.Request) string {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	var json map[string]interface{}
	var err = decoder.Decode(&json)
	if err != nil {
		return "not provided"
	} else {
		return json["request_id"].(string)
	}
}

func buildGophers() []Gopher {
	return []Gopher{
		{Id: rand.Int()}, {Id: rand.Int()}}
}

func Root(w http.ResponseWriter, r *http.Request) {
	var requestId = extractRequestId(r)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	rootResponse := &WombatResponse{
		RequestId: requestId,
		Gophers:   buildGophers()}

	response, _ := json.Marshal(rootResponse)
	io.WriteString(w, string(response))

	log.Printf("RequestId: %s\tPath: /", requestId)
}
