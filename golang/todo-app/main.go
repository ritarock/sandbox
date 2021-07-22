package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/users", readUsersAll)
	mux.HandleFunc("/users/create", createUsers)
	mux.HandleFunc("/users/read/", readUsers)
	mux.HandleFunc("/users/update/", updateUsers)
	mux.HandleFunc("/users/delete/", deleteUsers)

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
