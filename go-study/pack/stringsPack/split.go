package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Split("A,B,C,D,E", ","))
	fmt.Println(strings.SplitAfter("A,B,C,D,E", ","))
}
