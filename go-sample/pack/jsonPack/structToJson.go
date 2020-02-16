package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type User struct {
	Id      int
	Name    string
	Email   string
	Created time.Time
}

func main() {
	u := new(User)
	u.Id = 1
	u.Name = "hoge"
	u.Email = "hoge@example.com"
	u.Created = time.Now()

	bs, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bs))
}
