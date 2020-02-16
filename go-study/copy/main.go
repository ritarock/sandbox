package main

import (
	"io/ioutil"
	"os"
)

func main() {
	originName := os.Args[1]
	copyName := os.Args[2]

	f, err := os.Open(originName)
	defer f.Close()

	if err != nil {
		panic(err)
	}

	cf, err := os.Create(copyName)
	if err != nil {
		panic((err))
	}
	byteFile, _ := ioutil.ReadAll(f)
	cf.Write(byteFile)
}
