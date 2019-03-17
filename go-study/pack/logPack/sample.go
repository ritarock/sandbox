package main

import (
	"log"
)

func main() {
	log.Print("ログの1行目\n")
	log.Println("ログの2行目")
	log.Printf("ログの%d行目\n", 3)
}
