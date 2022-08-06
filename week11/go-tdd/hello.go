package main

import (
	"fmt"
)

const engGreeting = "Hello, "
const spaGreeting = "Hola, "
const freGreeting = "Bonjour, "

const spanish = "Spanish"
const french = "French"

func Hello(n string, lang string) string {
	if n == "" {
		n = "world"
	}

	return helloPrefix(lang) + n
}

func helloPrefix(lang string) (greet string) {
	switch lang {
	case spanish:
		greet = spaGreeting
	case french:
		greet = freGreeting
	default:
		greet = engGreeting
	}
	return
}

func main() {
	fmt.Println(Hello("Zack", spanish))
}
