package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/users", usersHandler)
	mux.HandleFunc("/user/", userHandler)
	mux.HandleFunc("/tasks", tasksHandler)
	mux.HandleFunc("/task", taskHandler)

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
