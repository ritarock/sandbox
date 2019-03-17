package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Id    int
	Name  string
	Email string
}

func main() {
	src := `
{
	"Id": 1,
	"Name": "hoge",
	"Email": "hoge@example.com"
}`

	u := new(User)
	err := json.Unmarshal([]byte(src), u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", u)
}
