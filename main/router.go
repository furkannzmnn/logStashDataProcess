package main

import (
	"fmt"
	"log"
	"net/http"
)

func Router() {
	fmt.Println("Server started on port 5690")
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleLargeRequest)
	log.Fatal(http.ListenAndServe(":5690", mux))

}
