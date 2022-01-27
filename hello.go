package main

import "fmt"

const englishHelloPrefix = "Hello, "
const welsh = "Welsh"
const welshHelloPrefix = "Helo, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	prefix = englishHelloPrefix

	if language == welsh {
		prefix = welshHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Oliver", ""))
}
