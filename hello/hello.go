package hello

import (
	"curiousjc/greetings"
	"log"
)

func Hello(name string) (realname string, string, bobname string) {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("we had a problem: ")
	log.SetFlags(0)

	message, err := greetings.Hello(name)

	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	bobname = "(and Bob)"
	suename := "(and Sue)"

	// If no error was returned, print the returned message
	// to the console.
	return message, suename, bobname
}
