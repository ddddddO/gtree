// Package http provides a set of HTTP Cloud Functions samples.
package http

import (
	"encoding/json"
	"net/http"
)

type Person struct {
	Name string `json: "name"`
	Age  int    `json: "age"`
}

// CORSEnabledFunction is an example of setting CORS headers.
// For more information about CORS and CORS preflight requests, see
// https://developer.mozilla.org/en-US/docs/Glossary/Preflight_request.
func CORSEnabledFunction(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "http://cors.ddddddo.work")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "http://cors.ddddddo.work")

	people := []Person{
		Person{
			Name: "ddd",
			Age:  27,
		},
		Person{
			Name: "fff",
			Age:  88,
		},
	}
	json.NewEncoder(w).Encode(people)
}
