package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const URL = "https://github.com/ritarock"

func main() {
	resp, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byteArray))
}
