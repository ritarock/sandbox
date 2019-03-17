package main

import (
	"fmt"
)

type User struct {
	Id    int
	Email string
}

func main() {
	u := &User{Id: 123, Email: "mail@exampl.com"}
	fmt.Printf("%v\n", u)
	fmt.Printf("%+v\n", u)
	fmt.Printf("%#v\n", u)
}
