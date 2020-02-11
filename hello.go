package main

import "fmt"

const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "

const french = "French"
const frenchHelloPrefix = "Bonjour, "

const german = "German"
const germanHelloPrefix = "Hallo, "

const englishHelloPrefix = "Hello, "

func Hello(name, language string) string {

	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case german:
		prefix = germanHelloPrefix
	}

	return prefix + name

}

func main() {
	fmt.Println(Hello("world", ""))
}
