package main

import "fmt"

func main() {
	hobies := []string{"baseball", "movies", "books"}
	fmt.Println(hobies)
	fmt.Println(hobies[0])
	newhobbies := hobies[1:]
	fmt.Println(newhobbies)
}
