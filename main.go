package main

import (
	"fmt"
	"github.com/joshcork/tdd-code-kata/greetings"
)

func main(){
	Start()
}

func Start() {

	message := GreetMe("World!")
	fmt.Println(message)

}

