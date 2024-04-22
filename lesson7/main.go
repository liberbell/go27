package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userName := make([]string, 2, 5)

	userName[0] = "Juliet"

	userName = append(userName, "Max")
	userName = append(userName, "Sam")
	userName = append(userName, "Sam")
	userName = append(userName, "Sam")
	userName = append(userName, "Sam")
	userName = append(userName, "Sam")
	fmt.Println(userName)

	courcesRaitings := make(floatMap, 3)
	courcesRaitings["go"] = 4.7
	courcesRaitings["react"] = 4.8
	courcesRaitings["angular"] = 4.5
	courcesRaitings["node"] = 4.9
	// fmt.Println(courcesRaitings)
	courcesRaitings.output()

	for index, value := range userName {
		fmt.Println(index, value)
	}

	for key, value := range courcesRaitings {
		fmt.Println(key, value)
	}
}
