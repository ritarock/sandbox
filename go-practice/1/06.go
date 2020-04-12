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

func makeUnion(x, y []string) []string {
	x = append(x, y...)
	m := make(map[string]bool)
	ansArr := make([]string, 0)

	for _, ele := range x {
		if !m[ele] {
			m[ele] = true
			ansArr = append(ansArr, ele)
		}
	}

	return ansArr
}

func makeIntersection(x, y []string) []string {
	m := make(map[string]bool)
	ansArr := make([]string, 0)

	for _, ele := range x {
		m[ele] = true
	}
	for _, ele := range y {
		if m[ele] {
			ansArr = append(ansArr, ele)
		}
	}

	return ansArr
}

func makeDifference(x, y []string) []string {
	m := make(map[string]bool)
	ansArr := make([]string, 0)

	for _, ele := range y {
		m[ele] = true
	}
	for _, ele := range x {
		if !m[ele] {
			ansArr = append(ansArr, ele)
		}
	}

	return ansArr
}

func main() {
	word1 := "paraparaparadise"
	word2 := "paragraph"

	x := nGram(word1, 2)
	y := nGram(word2, 2)

	fmt.Println(makeUnion(x, y))
	fmt.Println(makeIntersection(x, y))
	fmt.Println(makeDifference(x, y))
}