package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/json", responseHandler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func responseHandler(w http.ResponseWriter, r *http.Request) {
	response := []byte(`{"id": 1, "name": "id1"}`)
	defer func() {
		w.Header().Set("Content-Type", "aplication/json")
		fmt.Fprint(w, string(response))
	}()
}
