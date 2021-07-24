package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todo-app/data"

	"github.com/gorilla/mux"
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
	params := mux.Vars(request)
	userId, err := strconv.Atoi(params["user_id"])
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
	params := mux.Vars(request)
	userId, err := strconv.Atoi(params["user_id"])
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
	params := mux.Vars(request)
	userId, err := strconv.Atoi(params["user_id"])
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

func readTasksAll(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	userId, err := strconv.Atoi(params["user_id"])
	if err != nil {
		danger(err, "Cannot parse form")
		badRequest(writer, request, 400, "Bad Request")
		return
	}
	tasks := data.TasksAll(userId)
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

func createTasks(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		danger(err, "Cannot parse form")
		badRequest(writer, request, 400, "Bad Request")
		return
	}
	params := mux.Vars(request)
	userId, err := strconv.Atoi(params["user_id"])
	if err != nil {
		danger(err, "Cannot parse form")
		badRequest(writer, request, 400, "Bad Request")
		return
	}

	var task data.Task
	json.NewDecoder(request.Body).Decode(&task)
	task.UserId = userId
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

func readTasks(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		danger(err, "Cannot parse form")
		badRequest(writer, request, 400, "Bad Request")
		return
	}
	params := mux.Vars(request)
	userId, err := strconv.Atoi(params["user_id"])
	if err != nil {
		danger(err, "Cannot parse form")
		badRequest(writer, request, 400, "Bad Request")
		return
	}
	taskId, err := strconv.Atoi(params["task_id"])
	if err != nil {
		danger(err, "Cannot parse form")
		badRequest(writer, request, 400, "Bad Request")
		return
	}
	var task data.Task
	task.ID = taskId
	task.UserId = userId
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

func updateTasks(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		danger(err, "Cannot parse form")
		badRequest(writer, request, 400, "Bad Request")
		return
	}
	params := mux.Vars(request)
	userId, err := strconv.Atoi(params["user_id"])
	if err != nil {
		danger(err, "Cannot parse form")
		badRequest(writer, request, 400, "Bad Request")
		return
	}
	taskId, err := strconv.Atoi(params["task_id"])
	var task data.Task
	task.ID = taskId
	task.UserId = userId
	task.Read()
	var newTask data.Task
	json.NewDecoder(request.Body).Decode(&newTask)
	task.Update(newTask)
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

func deleteTasks(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		danger(err, "Cannot parse form")
		badRequest(writer, request, 400, "Bad Request")
		return
	}
	params := mux.Vars(request)
	userId, err := strconv.Atoi(params["user_id"])
	if err != nil {
		danger(err, "Cannot parse form")
		badRequest(writer, request, 400, "Bad Request")
		return
	}
	taskId, err := strconv.Atoi(params["task_id"])
	var task data.Task
	task.ID = taskId
	task.UserId = userId
	deleteTask := task
	deleteTask.Read()
	task.Delete()
	var response struct {
		Code int         `json:"code"`
		Data []data.Task `json:"data"`
	}
	response.Code = 200
	response.Data = append(response.Data, deleteTask)
	r, _ := json.Marshal(response)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(r)
}
