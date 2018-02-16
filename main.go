package main

import (
	"fmt"
	"errors"
)

func main() {
	var NegativeError = errors.New("This is an error Yo!")

	fmt.Print(NegativeError)
}