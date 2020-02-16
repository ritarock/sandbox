package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Replace("AAAAA", "A", "X", 2))
	fmt.Println(strings.Replace("AAAAA", "A", "X", -1))
}
