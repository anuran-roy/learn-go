package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Name is Empty!")
	}
	message := fmt.Sprintf("Hi %v. Welcome!", name)
	return message, nil
}
