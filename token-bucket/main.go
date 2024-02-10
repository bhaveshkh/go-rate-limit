package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

func endPointHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	// var message Message
	message := Message{
		Status: "Successfull",
		Body:   "Hi! You've reached the API. How may I you.",
	}

	err := json.NewEncoder(writer).Encode(&message)
	if err != nil {
		return
	}
}

func main() {
	http.Handle("/ping", rateLimiter(endPointHandler))
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Println("There was an error listening in port :8080", err)
	}

}
