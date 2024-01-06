package main

import (
	"curiousjc/goodbye"
	"curiousjc/greetings"
	"curiousjc/hello"
	"fmt"
	"log"
)

func main() {

	//using the hello module and just spitting out the multi-valued return from the exported "Hello" function
	fmt.Println(hello.Hello("CuriousJC"))

	//USING THE greetings MODULE directly
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned map of
	// messages to the console.
	fmt.Println(messages)

	//using the goodbye module's exported "Goodbye" function
	fmt.Println(goodbye.Goodbye())
}
