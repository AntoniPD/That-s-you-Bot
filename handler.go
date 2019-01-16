package main

import (
	"encoding/json"
	"msgbot/models/messenger"
	"net/http"
	"os"
)

// VerificationEndpoint -
// Handles get request from the messenger api for bot verification.
// Returns them an int which they require.
func VerificationEndpoint(w http.ResponseWriter, r *http.Request) {
	challenge := r.URL.Query().Get("hub.challenge")
	token := r.URL.Query().Get("hub.verify_token")

	if token == os.Getenv("VERIFY_TOKEN") {
		w.WriteHeader(200)
		w.Write([]byte(challenge))
	} else {
		w.WriteHeader(404)
		w.Write([]byte("Error, wrong validation token"))
	}
}

// MessagesEndpoint -
// Handles the messages send from the users.
func MessagesEndpoint(w http.ResponseWriter, r *http.Request) {
	var callback models.Callback
	json.NewDecoder(r.Body).Decode(&callback)
	if callback.Object == "page" {
		for _, entry := range callback.Entry {
			for _, event := range entry.Messaging {
				ProcessMessage(event)
			}
		}
		w.WriteHeader(200)
		w.Write([]byte("Got your message"))
	} else {
		w.WriteHeader(404)
		w.Write([]byte("Message not supported"))
	}
}
