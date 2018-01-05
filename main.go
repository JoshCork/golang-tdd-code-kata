package main

import (
	"fmt"
	"github.com/joshcork/tdd-code-kata/messages"
)

func main(){
	Start()
}

func Start() {

	message := messages.GreetMe("World!")
	fmt.Println(message)

}

