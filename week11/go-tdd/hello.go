package main

import (
	"fmt"
)

func Hello(n string) string {
	return "Hello, " + n
}

func main() {
	fmt.Println(Hello("Zack"))
}
