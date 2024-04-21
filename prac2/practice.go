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

	products := []Product{
		Product{
			"First product",
			"A First Product",
			10.99,
		},
		Product{
			"Second product",
			"A Second Product",
			24.99,
		},
	}
	fmt.Println(products)

	newProduct := Product{
		"the third product",
		"A third product",
		4.99,
	}

	products = append(products, newProduct)
	fmt.Println(products)
}
