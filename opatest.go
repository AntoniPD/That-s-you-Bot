package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mux"
)

func HomeEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from mlabouardy :)")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeEndpoint)
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
