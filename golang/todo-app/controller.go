package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"todo-app/data"
)

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "This is index page")
}

func readUsersAll(writer http.ResponseWriter, request *http.Request) {
	users := data.UsersAll()
	writer.Header().Set("Content-Type", "application/json")
	var response struct {
		Code int         `json:"code"`
		Data []data.User `json:"data"`
	}
	response.Code = 200
	response.Data = append(response.Data, users...)
	r, _ := json.Marshal(response)
	writer.Write(r)
}

func createUsers(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		waring(request.Method, "Method Not Allow")
		badRequest(writer, request, 405, "Method Not Allow")
		return
	}
	err := request.ParseForm()
	if err != nil {
		danger(err, "Cannot parse form")
		badRequest(writer, request, 400, "Bad Request")
		return
	}
	var user data.User
	json.NewDecoder(request.Body).Decode(&user)
	user.Create()
	var response struct {
		Code int         `json:"code"`
		Data []data.User `json:"data"`
	}
	response.Code = 200
	response.Data = append(response.Data, user)
	r, _ := json.Marshal(response)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(r)
}

func readUsers(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		waring(request.Method, "Method Not Allow")
		badRequest(writer, request, 405, "Method Not Allow")
		return
	}
	sub := strings.TrimPrefix(request.URL.Path, "/users/read/")
	userId, err := strconv.Atoi(sub)
	if err != nil {
		danger(err, "Cannot parse form")
		badRequest(writer, request, 400, "Bad Request")
		return
	}
	var user data.User
	user.ID = userId
	user.Read()
	var response struct {
		Code int         `json:"code"`
		Data []data.User `json:"data"`
	}
	response.Code = 200
	response.Data = append(response.Data, user)
	r, _ := json.Marshal(response)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(r)
}

func updateUsers(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "PUT" {
		waring(request.Method, "Method Not Allow")
		badRequest(writer, request, 405, "Method Not Allow")
		return
	}
	sub := strings.TrimPrefix(request.URL.Path, "/users/update/")
	userId, err := strconv.Atoi(sub)
	if err != nil {
		danger(err, "Cannot parse form")
		badRequest(writer, request, 400, "Bad Request")
		return
	}
	var user data.User
	user.ID = userId
	user.Read()
	var newUser data.User
	json.NewDecoder(request.Body).Decode(&newUser)
	user.Update(newUser)
	user.Read()
	var response struct {
		Code int         `json:"code"`
		Data []data.User `json:"data"`
	}
	response.Code = 200
	response.Data = append(response.Data, user)
	r, _ := json.Marshal(response)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(r)
}

func deleteUsers(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "DELETE" {
		waring(request.Method, "Method Not Allow")
		badRequest(writer, request, 405, "Method Not Allow")
		return
	}
	sub := strings.TrimPrefix(request.URL.Path, "/users/delete/")
	userId, err := strconv.Atoi(sub)
	if err != nil {
		danger(err, "Cannot parse form")
		badRequest(writer, request, 400, "Bad Request")
		return
	}
	var user data.User
	user.ID = userId
	deleteUser := user
	deleteUser.Read()
	user.Delete()
	var response struct {
		Code int         `json:"code"`
		Data []data.User `json:"data"`
	}
	response.Code = 200
	response.Data = append(response.Data, deleteUser)
	r, _ := json.Marshal(response)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(r)
}

func readTasksAll(writer http.ResponseWriter, request *http.Request, user_id int) {
	tasks := data.TasksAll(user_id)
	writer.Header().Set("Content-Type", "application/json")
	var response struct {
		Code int         `json:"code"`
		Data []data.Task `json:"data"`
	}
	response.Code = 200
	response.Data = append(response.Data, tasks...)
	r, _ := json.Marshal(response)
	writer.Write(r)
}

func createTasks(writer http.ResponseWriter, request *http.Request, user_id int) {
	if request.Method != "POST" {
		waring(request.Method, "Method Not Allow")
		badRequest(writer, request, 405, "Method Not Allow")
		return
	}
	err := request.ParseForm()
	if err != nil {
		danger(err, "Cannot parse form")
		badRequest(writer, request, 400, "Bad Request")
		return
	}
	var task data.Task
	json.NewDecoder(request.Body).Decode(&task)
	task.UserId = user_id
	task.Create()
	var response struct {
		Code int         `json:"code"`
		Data []data.Task `json:"data"`
	}
	response.Code = 200
	response.Data = append(response.Data, task)
	r, _ := json.Marshal(response)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(r)
}

func readTasks(writer http.ResponseWriter, request *http.Request, user_id int) {
	if request.Method != "GET" {
		waring(request.Method, "Method Not Allow")
		badRequest(writer, request, 405, "Method Not Allow")
		return
	}
	sub := strings.TrimPrefix(request.URL.Path, "/users/"+strconv.Itoa(user_id)+"/tasks/read/")
	taskId, err := strconv.Atoi(sub)
	if err != nil {
		danger(err, "Cannot parse form")
		badRequest(writer, request, 400, "Bad Request")
		return
	}
	var task data.Task
	task.ID = taskId
	task.UserId = user_id
	task.Read()
	var response struct {
		Code int         `json:"code"`
		Data []data.Task `json:"data"`
	}
	response.Code = 200
	response.Data = append(response.Data, task)
	r, _ := json.Marshal(response)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(r)
}

func updateTasks(writer http.ResponseWriter, request *http.Request, user_id int) {
}

func deleteTasks(writer http.ResponseWriter, request *http.Request, user_id int) {
}
