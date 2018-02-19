package main

import (
	"fmt"

	"github.com/joshcork/tdd-code-kata/calculator"
)

func main() {

	testing, err := calculator.Add("-1")
	if err != nil {
		fmt.Printf("\nwe expected an error, just not this one: %s\n\n", err)
		fmt.Print(err.Error())
	} else {

		fmt.Printf("everything's cool: %d", testing)
	}

}
