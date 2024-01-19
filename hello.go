package main

import "fmt"

const ukrainian = "Ukrainian"
const french = "French"
const englishHelloPrefix = "Hello, "
const ukrainianHelloPrefix = "Привіт, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := englishHelloPrefix
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case ukrainian:
		prefix = ukrainianHelloPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("World", "Ukrainian"))
}
