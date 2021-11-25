package greetings

import (
	"errors"
	"fmt"
	"time"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Name is Empty!")
	}
	message := fmt.Sprintf("%v, %v!", randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string] string, error){
	messages := make(map[string] string)

	for _, name:= range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func init(){
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi there",
		"Nice to meet you",
		"Great to see you",
	}

	return formats[rand.Intn(len(formats))]
}