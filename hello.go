package main

import (
	"fmt"
	"log"

	greetings "github.com/caraar12345/go-tutorial-greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
