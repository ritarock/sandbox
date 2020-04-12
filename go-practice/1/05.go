package main

import (
	"fmt"
)

func nGram(s string, n int) []string {
	rs := []rune(s)
	var response []string

	for i := 0; i < len(rs)-n+1; i++ {
		response = append(response, string(rs[i:i+n]))
	}
	return response
}

func main() {
	word := "n-gram"
	ans := nGram(word, 2)

	fmt.Println(ans)
}
