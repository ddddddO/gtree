package main

import (
	"log"
	"net/http"
)

func main() {
	// BasicAuth
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || user != "ddd" || pass != "ddd" {
			w.Header().Add("WWW-Authenticate", `Basic realm="local file directory"`)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("failed to auth"))
			return
		}

		w.Write([]byte(`Authenticated! Next, go to "/svg"`))
	})

	http.Handle("/svg", http.FileServer(http.Dir("../../_data")))

	// Simple static webserver:
	//log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("../../_data"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
