// package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"msgbot/models"
// 	"net/http"
// 	"os"
// 	"reflect"

// 	"github.com/mux"
// )

// const (
// 	FACEBOOK_API   = "https://graph.facebook.com/v2.6/me/messages?access_token=%s"
// 	IMAGE          = "http://37.media.tumblr.com/e705e901302b5925ffb2bcf3cacb5bcd/tumblr_n6vxziSQD11slv6upo3_500.gif"
// 	VISIT_SHOW_URL = "http://labouardy.com"
// )

// func main() {
// 	os.Setenv("PAGE_ACCESS_TOKEN", "EAAEvrRFgZCycBADN9EDxJF2Hu0gohDhCEAIll2W1uCCFDTpbLIDELKWLgEncr9p0DA1X4OKikC2bsqr12d0Pmk13w6WMueQQL1LYKWhXsjB5Swj45JdrXwu2IfCAArhkEuRFbRKUqi90UoYKq6hQ542v0qLnIot8xYYLbxQZDZD")
// 	r := mux.NewRouter()
// 	r.HandleFunc("/webhook", VerificationEndPoint).Methods("GET")
// 	r.HandleFunc("/webhook", MessagesEndPoint).Methods("POST")
// 	log.Fatal(http.ListenAndServe(":8080", r))
// }

// func VerificationEndPoint(w http.ResponseWriter, r *http.Request) {
// 	challenge := r.URL.Query().Get("hub.challenge")
// 	mode := r.URL.Query().Get("hub.mode")
// 	token := r.URL.Query().Get("hub.verify_token")
// 	if mode != "" && token == os.Getenv("VERIFY_TOKEN") {
// 		w.WriteHeader(200)
// 		w.Write([]byte(challenge))
// 	} else {
// 		w.WriteHeader(404)
// 		w.Write([]byte("Error, wrong validation token"))
// 	}
// }

// func MessagesEndPoint(w http.ResponseWriter, r *http.Request) {
// 	var callback models.Callback
// 	json.NewDecoder(r.Body).Decode(&callback)
// 	if callback.Object == "page" {
// 		for _, entry := range callback.Entry {
// 			for _, event := range entry.Messaging {
// 				if !reflect.DeepEqual(event.Message, models.Message{}) && event.Message.Text != "" {
// 					ProcessMessage(event)
// 				}
// 			}
// 		}
// 		w.WriteHeader(200)
// 		w.Write([]byte("Got your message"))
// 	} else {
// 		w.WriteHeader(404)
// 		w.Write([]byte("Message not supported"))
// 	}
// }

// func ProcessMessage(event models.Messaging) {
// 	client := &http.Client{}
// 	response := http.Response{
// 		Recipient: User{
// 			ID: event.Sender.ID,
// 		},
// 		Message: Message{
// 			Attachment: &Attachment{
// 				Type: "image",
// 				Payload: Payload{
// 					URL: IMAGE,
// 				},
// 			},
// 		},
// 	}

// 	body := new(bytes.Buffer)
// 	json.NewEncoder(body).Encode(&response)

// 	url := fmt.Sprintf(FACEBOOK_API, os.Getenv("PAGE_ACCESS_TOKEN"))
// 	req, err := http.NewRequest("POST", url, body)
// 	req.Header.Add("Content-Type", "application/json")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer resp.Body.Close()
// }
