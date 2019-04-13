package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", FormHandler)

	log.Println("Serving on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
