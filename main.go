package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Foo struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type Foos []Foo

func singleRecord(w http.ResponseWriter, r *http.Request) {
	foos := Foos{
		Foo{Name: "Test Guy Name", Id: "1234"},
	}

	fmt.Println("Single Record Endpoint Hit")
	json.NewEncoder(w).Encode(foos)
}

func handleRequests() {
	http.HandleFunc("/", singleRecord)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
