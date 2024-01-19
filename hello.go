package main

import "fmt"

const ukrainian = "Ukrainian"
const englishHelloPrefix = "Hello, "
const ukrainianHelloPrefix = "Привіт, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == ukrainian {
		return ukrainianHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("World", "Ukrainian"))
}
