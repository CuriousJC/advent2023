package main

import (
	"fmt"

	"github.com/curiousjc/golang-learning/pkg/chorus"
	"github.com/curiousjc/golang-learning/pkg/conversation"
	"github.com/curiousjc/golang-learning/pkg/goodbye"
	"github.com/curiousjc/golang-learning/pkg/random"
)

func main() {
	fmt.Println("Hello!!!")

	fmt.Println("Here comes a random number from our external package...")
	fmt.Println(random.Get())
	fmt.Println(goodbye.Goodbye())
	fmt.Println(conversation.Get())
	fmt.Println(chorus.Get())

}
