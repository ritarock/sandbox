package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := "file"

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Println(strings.Split(scanner.Text(), ","))
	}

	if err = scanner.Err(); err != nil {
		panic(err)
	}
}
