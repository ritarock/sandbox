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
	err := request.ParseForm()
	if err != nil {
		fmt.Println(err)
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
	sub := strings.TrimPrefix(request.URL.Path, "/users/read/")
	userId, err := strconv.Atoi(sub)
	if err != nil {
		fmt.Println(err)
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
	sub := strings.TrimPrefix(request.URL.Path, "/users/update/")
	userId, err := strconv.Atoi(sub)
	if err != nil {
		fmt.Println(err)
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
	sub := strings.TrimPrefix(request.URL.Path, "/users/delete/")
	userId, err := strconv.Atoi(sub)
	if err != nil {
		fmt.Println(err)
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
