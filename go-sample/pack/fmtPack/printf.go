package main

import (
	"fmt"
)

func main() {
	n := 100
	f := 3.14

	fmt.Printf("%d * %.2f = %.2f\n", n, f, float64(n)*f)
}
