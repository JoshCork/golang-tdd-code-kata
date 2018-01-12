package main

import (
	"fmt"
	"github.com/joshcork/tdd-code-kata/calculator"
)

func main() {
	myResult := calculator.Add("")
	fmt.Printf("The result is: %d", myResult)
}