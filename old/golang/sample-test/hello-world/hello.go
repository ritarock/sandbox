package main

import "fmt"

const Spanish = "Spanish"
const French = "French"
const EnglishHelloPrefix = "Hello, "
const SpanishHelloPrefix = "Hola, "
const FrenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(langage string) (prefix string) {
	switch langage {
	case French:
		prefix = FrenchHelloPrefix
	case Spanish:
		prefix = SpanishHelloPrefix
	default:
		prefix = EnglishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("", ""))
}
