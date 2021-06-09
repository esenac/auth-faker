package main

import (
	"encoding/json"
	"net/http"
)

// GetHandler handles the index route
func GetHandler(data interface{}) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		jsonBody, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Error converting results to json",
				http.StatusInternalServerError)
		}
		w.Write(jsonBody)
	}
}
