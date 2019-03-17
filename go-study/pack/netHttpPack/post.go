package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	vs := url.Values{}
	vs.Add("id", "1")
	fmt.Println(vs.Encode())

	res, err := http.PostForm("https://example.com/commements/post", vs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
