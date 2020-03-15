package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// ReadFile read file
func readFile(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	return string(data)
}

// writeFile write file
func writeFile(path string, data string) string {
	err := ioutil.WriteFile(path, []byte(data), 0755)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Sprintf("Create %s", path)
}

// removeFile remove file
func removeFile(path string) string {
	err := os.Remove(path)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Sprintf("Remove %s", path)
}


func main() {
	fmt.Println(readFile("./read.file"))
	fmt.Println(writeFile("./write.file", "hello world"))
	fmt.Println(removeFile("./write.file"))
}