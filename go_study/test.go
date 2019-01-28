package main

import (
	"fmt"
)

func main() {
	fruits := [3]string{"A", "B", "C"}

	for i, s := range fruits {
		fmt.Printf("fruits[%d]=%s\n", i, s)
	}
}
