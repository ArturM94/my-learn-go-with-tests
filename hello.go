package main

import "fmt"

const helloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		return helloPrefix + "World"
	}
	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("World"))
}
