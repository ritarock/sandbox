package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "パタトクカシーー"
	arr := strings.Split(word, "")
	var ans string

	for i, s := range arr {
		if i%2 == 0 {
			ans = ans + s
		}
	}

	fmt.Println(ans)
}
