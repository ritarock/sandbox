package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var mux *http.ServeMux

func main() {
	mux = http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/users", readUsersAll)
	mux.HandleFunc("/users/create", createUsers)
	mux.HandleFunc("/users/read/", readUsers)
	mux.HandleFunc("/users/update/", updateUsers)
	mux.HandleFunc("/users/delete/", deleteUsers)
	mux.HandleFunc("/users/", taskRouter)

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

func taskRouter(writer http.ResponseWriter, request *http.Request) {
	sub := strings.TrimPrefix(request.URL.Path, "/users/")
	splitPath := strings.Split(sub, "/")
	user_id, err := strconv.Atoi(splitPath[0])
	if err != nil {
		fmt.Println(err)
	}
	if len(splitPath) > 2 {
		method := splitPath[2]
		switch method {
		case "create":
			createTasks(writer, request, user_id)
		case "read":
			readTasks(writer, request, user_id)
		case "update":
			updateTasks(writer, request, user_id)
		case "delete":
			deleteTasks(writer, request, user_id)
		default:
			fmt.Fprintf(writer, "err")
		}
	} else {
		readTasksAll(writer, request, user_id)
	}
}
