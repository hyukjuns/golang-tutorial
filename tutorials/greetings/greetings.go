package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randFormat(), name)
	// message := fmt.Sprintf(randFormat())
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	// make(map[key-type]value-type)
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func randFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
