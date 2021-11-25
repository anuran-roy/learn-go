package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Get a greeting message and print it.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := greetings.Hello("Anuran")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	messages, err := greetings.Hellos([]string{"Anuran", "Akshit", "Ayan", "Aayush"})

	if err != nil {
		log.Fatal(err)
	}
	for _, msg := range messages{
		fmt.Println(msg)
	}
}
