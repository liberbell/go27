package main

import "fmt"

func main() {
	hobies := []string{"baseball", "movies", "books"}
	fmt.Println(hobies)
	fmt.Println(hobies[0])
	newhobbies := hobies[:2]
	fmt.Println(newhobbies)

	fmt.Println(cap(newhobbies))
	newhobbies = newhobbies[1:3]
	fmt.Println(newhobbies)

	courceGoals := []string{"Learn go", "Learn all the basics"}
	fmt.Println(courceGoals)
}
