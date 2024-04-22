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
}
