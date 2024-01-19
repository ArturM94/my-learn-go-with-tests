package main

import "fmt"

const ukrainian = "Ukrainian"
const french = "French"
const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const ukrainianHelloPrefix = "Привіт, "
const frenchHelloPrefix = "Bonjour, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := englishHelloPrefix
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case ukrainian:
		prefix = ukrainianHelloPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("World", "Ukrainian"))
}
