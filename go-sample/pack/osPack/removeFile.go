package main

import (
	"log"
	"os"
)

func main() {
	err := os.Remove("foo.txt")
	if err != nil {
		log.Fatal(err)
	}
}
