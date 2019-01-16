package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mux"
)

func main() {
	initEnv()
	usr := GetUserInfo("2170140013031793")
	fmt.Println(usr)
	r := mux.NewRouter()
	r.HandleFunc("/webhook", VerificationEndpoint).Methods("GET")
	r.HandleFunc("/webhook", MessagesEndpoint).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}
