package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/jahongir", apiHandler)
	http.ListenAndServe(":8080", nil)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGet(w, r)
	case http.MethodPost:
		handlePost(w, r)
	case http.MethodPut:
		handlePut(w, r)
	case http.MethodDelete:
		handleDelete(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("id")
	if id != "" {
		fmt.Fprintf(w, "GET request received with ID: %s\n", id)
	} else {
		fmt.Fprint(w, "GET request received with no ID\n")
	}
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	response := fmt.Sprintf("POST request received with data: %v\n", data)
	fmt.Fprint(w, response)
}

func handlePut(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	response := fmt.Sprintf("PUT request received with data: %v\n", data)
	fmt.Fprint(w, response)
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("id")
	if id != "" {
		fmt.Fprintf(w, "DELETE request received for ID: %s\n", id)
	} else {
		fmt.Fprint(w, "DELETE request received with no ID\n")
	}
}
