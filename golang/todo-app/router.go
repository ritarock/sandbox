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
	switch request.Method {
	case "GET":
		fmt.Fprintf(writer, fmt.Sprintf("request: %v", request.Method))
	case "POST":
		err := request.ParseForm()
		if err != nil {
			fmt.Println(err)
		}
		var user User
		fmt.Println(request.Body)
		json.NewDecoder(request.Body).Decode(&user)
		db := gormConnect()
		db.Create(&user)
	}
}
