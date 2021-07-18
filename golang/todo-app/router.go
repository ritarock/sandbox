package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
)

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "this is index page")
}

func usersHandler(writer http.ResponseWriter, request *http.Request) {
	db := gormConnect()
	writer.Header().Set("Content-Type", "application/json")
	var response struct {
		Status int    `json:"status"`
		Data   []User `json:"data"`
	}
	switch request.Method {
	case "GET":
		var users []User
		db.Find(&users)
		_, err := json.Marshal(&users)
		if err != nil {
			fmt.Println(err)
			return
		}
		response.Status = 200
		response.Data = append(response.Data, users...)
		r, _ := json.Marshal(response)
		writer.Write(r)
	case "POST":
		err := request.ParseForm()
		if err != nil {
			fmt.Println(err)
			return
		}
		var user User
		json.NewDecoder(request.Body).Decode(&user)
		db.Create(&user)
		response.Status = 200
		response.Data = append(response.Data, user)
		r, _ := json.Marshal(response)
		writer.Write(r)
	default:
		fmt.Fprintf(writer, "Bad request")
	}
}

func userHandler(writer http.ResponseWriter, request *http.Request) {
	sub := strings.TrimPrefix(request.URL.Path, "/user")
	_, id := filepath.Split(sub)
	if id != "" {
		db := gormConnect()
		writer.Header().Set("Content-Type", "application/json")
		var response struct {
			Status int    `json:"status"`
			Data   []User `json:"data"`
		}
		switch request.Method {
		case "GET":
			var user User
			db.Find(&user, id)
			_, err := json.Marshal(&user)
			if err != nil {
				fmt.Println(err)
				return
			}
			response.Status = 200
			response.Data = append(response.Data, user)
			r, _ := json.Marshal(response)
			writer.Write(r)
		case "PUT":
			var user User
			var newUser User
			json.NewDecoder(request.Body).Decode(&newUser)
			db.Find(&user, id)
			if newUser.Name != "" {
				user.Name = newUser.Name
			}
			if newUser.Email != "" {
				user.Email = newUser.Email
			}
			db.Save(&user)
			response.Status = 200
			response.Data = append(response.Data, user)
			r, _ := json.Marshal(response)
			writer.Write(r)
		case "DELETE":
			var user User
			db.Find(&user, id)
			db.Delete(&user)
			response.Status = 200
			response.Data = append(response.Data, user)
			r, _ := json.Marshal(response)
			writer.Write(r)
		default:
			fmt.Fprintf(writer, "Bad request")
		}
	} else {
		fmt.Fprintf(writer, "Bad request")
	}
}

func tasksHandler(writer http.ResponseWriter, r *http.Request) {
}

func taskHandler(writer http.ResponseWriter, r *http.Request) {
}
