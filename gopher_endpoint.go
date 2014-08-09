package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type WombatResponse struct {
	RequestId string   `json:"request_id"`
	Gophers   []Gopher `json:"gophers"`
}

type Gopher struct {
	Id int `json:"id"`
}

func main() {
	log.Printf("Starting gopher_endpoint...")

	http.HandleFunc("/", Root)

	err := http.ListenAndServe(":9292", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func Root(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	gophers := []Gopher{
		{Id: 123}, {Id: 456}}

	rootResponse := &WombatResponse{
		RequestId: "123f",
		Gophers:   gophers}

	response, _ := json.Marshal(rootResponse)
	io.WriteString(w, string(response))

	log.Printf("/")
}
