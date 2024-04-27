package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Holla, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name, lang string) string {

	if name == "" {
		name = "world"
	}

	return greetingPrefix(lang) + name

}

func main() {
	fmt.Println(Hello("world", ""))
}

func greetingPrefix(l string) string {
	switch l {
	case spanish:
		return spanishHelloPrefix
	case french:
		return frenchHelloPrefix
	default:
		return englishHelloPrefix
	}
}
