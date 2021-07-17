package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "this is index page")
}

func userHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		err := request.ParseForm()
		if err != nil {
			fmt.Println(err)
		}
		var user User
		json.NewDecoder(request.Body).Decode(&user)
		db := gormConnect()
		db.Create(&user)

	} else {
		fmt.Fprintf(writer, "Bad Request")
	}
}

func usersHandler(writer http.ResponseWriter, request *http.Request) {
	var users []User
	db := gormConnect()
	db.Find(&users)
	response, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(response)
}
