package main

import "fmt"

func cipher(s string) []rune {
	rs := []rune(s)
	var u []rune

	for _, s := range rs {
		if s >= []rune("a")[0] && s <= []rune("z")[0] {
			u = append(u, 219 - s)
		} else {
			u = append(u, s)
		}
	}
	return u
}

func main() {
	word := "This is a pen."

	fmt.Println(string(cipher(word)))
	fmt.Println(string(cipher(string(cipher(word)))))
}