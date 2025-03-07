package greetings

import (
	"fmt"
	"errors"
)

// Hello returns a greeting for the named person.

func Hello (name string) (string, error) {
	// If no name was given, return an error with message.
	if name == "" {
		return "", errors.New("empty name")
	}	

	//return a greeting that embeds the name in message
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}