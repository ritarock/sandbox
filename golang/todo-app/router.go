package main

import (
	"fmt"
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "This is index page")
}
