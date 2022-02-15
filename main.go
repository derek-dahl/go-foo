package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Foo Struct (Model)
type Foo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Foos []Foo

var fooCollection []Foo

func deleteFoo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test DELETE endpoint hit")
}

// Create single Foo
func createFoo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var foo Foo
	_ = json.NewDecoder(r.Body).Decode(&foo)
	uuid := uuid.New()
	foo.Id = uuid.String()
	foos = append(foos, foo)
	json.NewEncoder(w).Encode(foo)
}

// Get single Foo
func getFoo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Params Variable
	params := mux.Vars(r) // Get params
	//Loop through foos to find id
	for _, item := range foos {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Foo{})
}

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
	router.HandleFunc("/foo", createFoo).Methods("POST")
	router.HandleFunc("/foo/{id}", deleteFoo).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
