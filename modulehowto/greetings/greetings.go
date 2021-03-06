package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	var message string = fmt.Sprintf("Hi, %v. Welcome!", name)

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

// lowercase == private, also init() is called automatically
func init() {
	rand.Seed(time.Now().UnixNano())
}

// uppercase == public
func RandomFormat() string {
	formats := []string{
		"Hi, %v, welcome",
		"Great to see you, %v",
		"Hail %v!!!",
	}
	return formats[rand.Intn(len(formats))]
}
