package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "

func main() {
	fmt.Println(Hello("Potato"))
}

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}
