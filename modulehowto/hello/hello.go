package main

import (
	"fmt"
	"log"
	"modulehowto/greetings"
	"os"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("no input supplied")
	}
	var name string
	name = args[0]
	if name == "" {
		log.Fatal("empty name")
	}

	var response string
	var err error

	if len(args) == 2 && args[1] == "--random" {
		response = fmt.Sprintf(greetings.RandomFormat(), name)
		fmt.Println(response)
	} else if len(args) == 1 {
		response, err = greetings.Hello(name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(response)
	} else {
		responses, err := greetings.Hellos(args)
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range responses {
			fmt.Println(v)
		}
	}
}
