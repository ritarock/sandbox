package main

import (
	"fmt"
	"strings"
)

func getRune(s string, i int) string {
	rs := []rune(s)
	return string(rs[i])
}

func main() {
	sentence := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	var ans string

	for _, s := range strings.Split(sentence, " ") {
		ans = ans + getRune(s, 0)
	}

	fmt.Println(ans)
}
