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
	dones := make([]chan bool, 4)
	// done := make(chan bool)

	dones[0] = make(chan bool)
	go greet("Nice to meet you", dones[0])
	dones[1] = make(chan bool)
	go greet("How are you", dones[1])
	go slowGreet("How ... are ... you ...", done)
	go greet("I home you`re liking the course", done)
	<-done
	<-done
	<-done
	<-done
}
