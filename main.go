package main

import (
	"encoding/json"
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

var foos []Foo

// Delete Single Foo
func deleteFoo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	for index, item := range foos {
		if item.Id == params["id"] {
			w.WriteHeader(http.StatusNoContent)
			foos = append(foos[:index], foos[index+1:]...)
			break
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

// Create Single Foo
func createFoo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var foo Foo
	_ = json.NewDecoder(r.Body).Decode(&foo)
	uuid := uuid.New()
	foo.Id = uuid.String()
	foos = append(foos, foo)
	json.NewEncoder(w).Encode(foo)
}

// Update Single Foo
func updateFoo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	for index, item := range foos {
		if item.Id == params["id"] {
			foos = append(foos[:index], foos[index+1:]...)
			var foo Foo
			_ = json.NewDecoder(r.Body).Decode(&foo)
			foo.Id = params["id"]
			foos = append(foos, foo)
			json.NewEncoder(w).Encode(foo)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

// Get Single Foo
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
	w.WriteHeader(http.StatusNotFound)
}

// Get All Foos
func getFoos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foos)
}

func main() {
	//Init Router
	router := mux.NewRouter().StrictSlash(true)

	// Route Handlers / Endpoints
	router.HandleFunc("/foo/{id}", getFoo).Methods("GET")
	router.HandleFunc("/foos", getFoos).Methods("GET")
	router.HandleFunc("/foo", createFoo).Methods("POST")
	router.HandleFunc("/foo/{id}", updateFoo).Methods("PUT")
	router.HandleFunc("/foo/{id}", deleteFoo).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
