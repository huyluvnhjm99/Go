package main

import "fmt"

type bot interface {
	greeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	greeting(eb)
	greeting(sb)
}

func greeting(b bot) {
	fmt.Println(b.greeting())
}

func (englishBot) greeting() string {
	return "Hello there!"
}

func (spanishBot) greeting() string {
	return "Hola!"
}
