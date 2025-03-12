package greetings

import (
	"fmt"
	"errors"
	"math/rand"
)

// Hello returns a greeting for the named person.

func Hello (name string) (string, error) {
	// If no name was given, return an error with message.
	if name == "" {
		return "", errors.New("empty name")
	}	

	//return a greeting that embeds the name in message
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates names to greeting messages

func Hellos (names []string) (map[string]string, error) {
	// A map to associate names with messages
	messages := make(map[string]string)
	// Loop Through the received slice of names, calling
	// the Hello function
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
} 

// randomFormat Returns one of a set of greeting messages. The returned
// message is selected at random

func randomFormat() string {
	// A slice of message formats
	formats:= []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return the randomly selected message format by specifying a random index
	// a random index for the slice formats
	
	return formats [rand.Intn(len(formats))]
}