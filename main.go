package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Project struct {
	id    string `json:"id"`
	title string `json:"title"`
	link  string `json:link`
}

func main() {
	// Let's build a basic router
	router := mux.NewRouter()

	// Build the routes
	router.HandleFunc("/api", sendHello).Methods("GET")
	router.HandleFunc("/api/health", checkHealth).Methods("GET")

	// Let's listen to the port for an endpoint
	log.Fatal(http.ListenAndServe(":8000", router))
}

func sendHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Hello World")
}

func checkHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Health seems fine")
}
