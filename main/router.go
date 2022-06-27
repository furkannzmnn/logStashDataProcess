package main

import (
	"log"
	"net/http"
)

func Router() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleLargeRequest)
	log.Fatal(http.ListenAndServe(":5690", mux))

}
