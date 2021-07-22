package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

var logger *log.Logger

func init() {
	file, err := os.OpenFile("log/todo-app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Faild to open log file", err)
	}
	logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

func badRequest(writer http.ResponseWriter, request *http.Request, code int, message string) {
	writer.Header().Set("Content-Type", "application/json")
	var response struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	response.Code = code
	response.Message = message
	r, _ := json.Marshal(response)
	writer.Write(r)
}

// for logging
func info(args ...interface{}) {
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

func waring(args ...interface{}) {
	logger.SetPrefix("WARNING ")
	logger.Println(args...)
}

func danger(args ...interface{}) {
	logger.SetPrefix("DANGER ")
	logger.Println(args...)
}
