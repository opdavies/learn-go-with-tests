package main

import "fmt"

const englishHelloPrefix = "Hello, "
const welsh = "Welsh"
const welshHelloPrefix = "Helo, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == welsh {
		return welshHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Oliver", ""))
}
