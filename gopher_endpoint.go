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

func extractRequestId(r *http.Request) string {
  defer r.Body.Close()

  decoder := json.NewDecoder(r.Body)
  var json map[string]interface{}
  var err = decoder.Decode(&json)
  if err != nil {
    return "1234"
  } else {
    return json["request_id"].(string)
  }
}

func Root(w http.ResponseWriter, r *http.Request) {
  var requestId = extractRequestId(r)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	gophers := []Gopher{
		{Id: 123}, {Id: 456}}

	rootResponse := &WombatResponse{
		RequestId: requestId,
		Gophers:   gophers}

	response, _ := json.Marshal(rootResponse)
	io.WriteString(w, string(response))

  log.Printf("RequestId: %s\tPath: /", requestId)
}
