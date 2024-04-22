package learning

import (
	"fmt"
	"github.com/curiousjc/golang-learning/pkg/random"
	"github.com/curiousjc/golang-learning/pkg/version"
)

func Print() {
	fmt.Println("Hello From the golang-learning package!!!")
	fmt.Println("The referenced package is: ", version.Get())

	fmt.Println("Here comes a random number from our external package...")
	fmt.Println(random.Get())
}
