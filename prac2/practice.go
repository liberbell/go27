package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

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

	courceGoals[1] = "Learn all the details"
	courceGoals = append(courceGoals, "Learn all the basics")
	fmt.Println(courceGoals)
}
