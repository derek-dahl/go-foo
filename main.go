package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Foo struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type Foos []Foo

var fooCollection []Foo

func deleteFoo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test DELETE endpoint hit")
}

func postFoo(w http.ResponseWriter, r *http.Request) {
	//Generate uuid here
	uuid := uuid.New
	fmt.Fprintf(w, "Test POST endpoint hit", uuid)
}

func getFoo(w http.ResponseWriter, r *http.Request) {
	}

	fmt.Println("Single Record Endpoint Hit")
func getFoos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foos)
}

func main() {
	//Init router
	router := mux.NewRouter().StrictSlash(true)

	// Mock Data - @todo - implement DB
	foos = append(foos, Foo{Id: "1", Name: "Bob"})
	foos = append(foos, Foo{Id: "2", Name: "Steve"})

	// Route Handlers / Endpoints
	router.HandleFunc("/foo/{id}", getFoo).Methods("GET")
	router.HandleFunc("/foos", getFoos).Methods("GET")
	router.HandleFunc("/foos", createFoo).Methods("POST")
	router.HandleFunc("/foo/{id}", deleteFoo).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
