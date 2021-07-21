package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todo-app/data"
)

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "This is index page")
}

func getUsers(writer http.ResponseWriter, request *http.Request) {
	users, err := data.GetUsers()
	if err != nil {
		fmt.Println(err)
	}
	var response struct {
		Code int         `json:"code"`
		Data []data.User `json:"data"`
	}
	response.Code = 200
	response.Data = append(response.Data, users...)
	r, _ := json.Marshal(response)
	writer.Header().Set("Content-Type", "application/json")
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
