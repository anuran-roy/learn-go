package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	// Get a greeting message and print it.
	var message = greetings.Hello("Anuran")
	fmt.Println(message)
}
