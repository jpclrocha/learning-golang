package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("name cannot be empty")
	}
	message := fmt.Sprintf(randomGreet(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
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

func randomGreet() string {
	formats := []string{
		"Hi, %v, Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Weel met!",
	}
	return formats[rand.Intn(len(formats))]
}
