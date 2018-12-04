package main

import (
	"fmt"
)

func main() {
	var n, m int = 1, 10
	fmt.Println(sumofInt(n, m))
	fmt.Println(sumoSquare(n, m))
	fmt.Println(sumofCube(n, m))
}

func sumofInt(n, m int) int {
	a := 0
	for ; n <= m; n++ {
		a = a + n
	}
	return a
}

func sumoSquare(n, m int) int {
	a := 0
	for ; n <= m; n++ {
		a = n * n
	}
	return a
}

func sumofCube(n, m int) int {
	a := 0
	for ; n <= m; n++ {
		a = a + n*n*n
	}
	return a
}
