package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func shuffleString(s string) []string {
	shuffleWord := strings.Split(s, "'")
	rand.Seed(time.Now().UnixNano())
	for i := range shuffleWord {
		j := rand.Intn(i + 1)
		shuffleWord[i], shuffleWord[j] = shuffleWord[j], shuffleWord[i]
	}

	return shuffleWord
}

func typoglycemia(s string) string {
	var ans []string
	for _, ele := range strings.Split(s, " ") {
		if len(ele) >= 4 {
			firstWord := ele[0]
			lastWord := ele[len(ele)-1]
			tmpWord := shuffleString(ele[1 : len(ele)-1])

			ans = append(ans, string(firstWord)+string(lastWord)+strings.Join(tmpWord, ""))
		} else {
			ans = append(ans, ele)
		}
	}

	return strings.Join(ans, " ")
}

func main() {
	sentence := "I couldn’t believe that I could actually understand what I was reading : the phenomenal power of the human mind ."
	fmt.Println(typoglycemia(sentence))
}
