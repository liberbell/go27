package main

import "fmt"

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

	courcesRaitings := make(map[string]float64, 3)
	courcesRaitings["go"] = 4.7
	courcesRaitings["react"] = 4.8
	courcesRaitings["angular"] = 4.5
	courcesRaitings["node"] = 4.9
	fmt.Println(courcesRaitings)
}
