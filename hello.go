package main

import "fmt"

const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "

const french = "French"
const frenchHelloPrefix = "Bonjour, "

const englishHelloPrefix = "Hello, "

func Hello(name, language string) string {

	if name == "" {
		name = "World"
	}

	if language == spanish {
		return spanishHelloPrefix + name
	} else if language == french {
		return frenchHelloPrefix + name
	} else {
		return englishHelloPrefix + name
	}

}

func main() {
	fmt.Println(Hello("world", ""))
}
