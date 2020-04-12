package main

import (
	"fmt"
	"strings"
)

func main() {
	word1 := "パトカー"
	word2 := "タクシー"
	var ans []string

	arr1 := strings.Split(word1, "")
	arr2 := strings.Split(word2, "")

	for i := 0; i < len(arr1); i++ {
		ans = append(ans, arr1[i]+arr2[i])
	}

	fmt.Println(strings.Join(ans, ""))
}
