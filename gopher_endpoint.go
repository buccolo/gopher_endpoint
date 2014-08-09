package main

import (
	"io"
	"log"
	"net/http"
)

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

	io.WriteString(w, "{ \"gophers\": 0 }\n")

	log.Printf("/")
}
