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

	courcesRaitings := map[string]float64{}
	courcesRaitings["go"] = 4.7
	courcesRaitings["react"] = 4.8
	fmt.Println(courcesRaitings)
}
