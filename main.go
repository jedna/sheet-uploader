package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", FormHandler)

	log.Println("Serving on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
