package main

import (
	"fmt"

	"rsc.io/quote"

	"example.com/greetings"
)

func main() {
	fmt.Println("Hello World!")
	message := greetings.Hello("Anuran")
	fmt.Println(quote.Go())
}
