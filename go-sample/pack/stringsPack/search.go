package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Index("ABCDE", "A"))
	fmt.Println(strings.Index("ABCDE", "BCD"))
	fmt.Println(strings.Index("ABCDE", "X"))
}
