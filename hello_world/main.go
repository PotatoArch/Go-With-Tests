package main

import (
	"fmt"
)

func main() {
	fmt.Println(Hello("Potato"))
}

func Hello(name string) string {
	return "Hello, " + name
}
