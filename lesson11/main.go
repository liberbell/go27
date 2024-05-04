package main

import "fmt"

func greet(phrase string) {
	fmt.Println("Hello ", phrase)
}

func main() {
	greet("Nice to meet you")
	greet("How are you")
	slowGreet("How ... are ... you ...")
}
