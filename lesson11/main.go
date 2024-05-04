package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello ", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello ", phrase)
	doneChan <- true
}

func main() {
	done := make(chan bool)
	go greet("Nice to meet you", done)
	go greet("How are you", done)
	go slowGreet("How ... are ... you ...", done)
	go greet("I home you`re liking the course", done)
	<-done
}
