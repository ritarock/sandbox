package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	router.HandleFunc("/backend/tasks", getTasks).Methods("GET")
	router.HandleFunc("/backend/tasks", createTask).Methods("POST")
	router.HandleFunc("/backend/tasks/{task_id}", getTask).Methods("GET")
	router.HandleFunc("/backend/tasks/{task_id}", updateTask).Methods("POST")
	router.HandleFunc("/backend/tasks/{task_id}", deleteTask).Methods("DELETE")

	router.HandleFunc("/backend/tasks/{task_id}/comments", getComments).Methods("GET")
	router.HandleFunc("/backend/tasks/{task_id}/comments", createComment).Methods("POST")
	router.HandleFunc("/backend/tasks/{task_id}/comments/{comment_id}", getComment).Methods("GET")
	router.HandleFunc("/backend/tasks/{task_id}/comments/{comment_id}", updateComment).Methods("POST")
	router.HandleFunc("/backend/tasks/{task_id}/comments/{comment_id}", deleteComment).Methods("DELETE")

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: router,
	}
	server.ListenAndServe()
}
