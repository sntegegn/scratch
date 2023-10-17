package main

import (
	"fmt"
)

const (
	french  = "French"
	spanish = "Spanish"

	englishHelloPrefix = "Hello "
	spanishHelloPrefix = "Hola "
	frenchHelloPrefix  = "Bonjour "
)

func Hello(name, lang string) string {
	if name == "" {
		name = "world"
	}

	return greetingWithPrefix(lang) + name
}

func greetingWithPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

func main() {
	fmt.Println(Hello("", ""))
}
