package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/foo", getFoo).Methods("GET")
	router.HandleFunc("/foo", postFoo).Methods("POST")
	router.HandleFunc("/foo", deleteFoo).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	handleRequests()
}
