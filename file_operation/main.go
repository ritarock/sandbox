package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// FileOpration path string
type FileOpration struct {
	path string
}

func (f FileOpration) writeFile(data string) string {
	err := ioutil.WriteFile(f.path, []byte(data), 0755)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Sprintf("Create %s", f.path)
}

func (f FileOpration) readFile() string {
	data, err := ioutil.ReadFile(f.path)
	if err != nil {
		fmt.Println(err)
	}
	return string(data)
}

func (f FileOpration) removeFile() string {
	err := os.Remove(f.path)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Sprintf("Remove %s", f.path)
}

func main() {
	wf := FileOpration{"./file.txt"}
	fmt.Println(wf.writeFile("hello"))
	fmt.Println(wf.readFile())
	fmt.Println(wf.removeFile())
}
