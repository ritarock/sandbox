package main

import "fmt"

func main() {
	var n int = 12
	primeFactor(n)
}

func primeFactor(n int) {
	for ; n%2 == 0; n /= 2 {
		fmt.Print(2, " ")
	}

	for i := 3; i*i <= n; i += 2 {
		for ; n%i == 0; n /= i {
			fmt.Print(i, " ")
		}
	}

	if n > 1 {
		fmt.Print(n)
	}
}
