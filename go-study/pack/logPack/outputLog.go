package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Create("test.log")
	if err != nil {
		return
	}

	log.SetOutput(f)
	log.Println("ログのメッセージ")
}
