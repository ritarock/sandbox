package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/users", readUsersAll).Methods("GET")
	mux.HandleFunc("/users", createUsers).Methods("POST")
	mux.HandleFunc("/users/{user_id}", readUsers).Methods("GET")
	mux.HandleFunc("/users/{user_id}", updateUsers).Methods("PUT")
	mux.HandleFunc("/users/{user_id}", deleteUsers).Methods("DELETE")

	mux.HandleFunc("/users/{user_id}/tasks", readTasksAll).Methods("GET")
	mux.HandleFunc("/users/{user_id}/tasks", createTasks).Methods("POST")
	mux.HandleFunc("/users/{user_id}/tasks/{task_id}", readTasks).Methods("GET")
	mux.HandleFunc("/users/{user_id}/tasks/{task_id}", updateTasks).Methods("PUT")
	mux.HandleFunc("/users/{user_id}/tasks/{task_id}", deleteTasks).Methods("DELETE")

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
