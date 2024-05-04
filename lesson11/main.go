package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello ", phrase)
}

func slowGreet(phrase string) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello ", phrase)
}

func main() {
	// greet("Nice to meet you")
	// greet("How are you")
	go slowGreet("How ... are ... you ...")
	// greet("I home you`re liking the course")
}
