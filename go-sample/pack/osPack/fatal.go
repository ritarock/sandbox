package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("foo")
	if err != nil {
		log.Fatal(err)
	}
}
