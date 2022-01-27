package main

import "fmt"

const englishHelloPrefix = "Hello, "
const welsh = "Welsh"
const welshHelloPrefix = "Helo, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	if language == welsh {
		prefix = welshHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Oliver", ""))
}
