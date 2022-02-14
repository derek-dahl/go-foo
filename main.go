package main

import (
	"fmt"
	"log"
	"net/http"
)

func singleRecord(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Single Record Endpoint Hit")
}

func handleRequests() {
	http.HandleFunc("/", singleRecord)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
