package main

import (
	"fmt"
	"strings"
)

func getRune(s string, i int) string {
	rs := []rune(s)
	return string(rs[:i])
}

func main() {
	sentence := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
	ans := map[int]string{}

	for i, s := range strings.Split(sentence, " ") {
		switch {
		case i+1 == 1:
			ans[i+1] = getRune(s, 1)
		case i+1 == 5:
			ans[i+1] = getRune(s, 1)
		case i+1 == 6:
			ans[i+1] = getRune(s, 1)
		case i+1 == 7:
			ans[i+1] = getRune(s, 1)
		case i+1 == 8:
			ans[i+1] = getRune(s, 1)
		case i+1 == 9:
			ans[i+1] = getRune(s, 1)
		case i+1 == 15:
			ans[i+1] = getRune(s, 1)
		case i+1 == 16:
			ans[i+1] = getRune(s, 1)
		case i+1 == 19:
			ans[i+1] = getRune(s, 1)
		default:
			ans[i+1] = getRune(s, 2)
		}
	}

	fmt.Println(ans)
}