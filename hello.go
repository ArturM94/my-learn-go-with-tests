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
	if language == ukrainian {
		return ukrainianHelloPrefix + name
	}
	if language == french {
		return frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("World", "Ukrainian"))
}
