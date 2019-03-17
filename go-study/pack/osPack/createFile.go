package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Create("foo.txt")
	fi, _ := f.Stat()

	fmt.Println(fi.Name())
	fmt.Println(fi.Size())
	fmt.Println(fi.IsDir())

	f.Write([]byte("Hello, World\n"))
	f.WriteAt([]byte("Golang"), 7)
	f.Seek(0, os.SEEK_END)
	f.WriteString("Yeah")
}
