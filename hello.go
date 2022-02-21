package main

import (
	"fmt"
	"log"

	greetings "github.com/caraar12345/go-tutorial-greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Aaron", "Jack", "Human"}

	message, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
