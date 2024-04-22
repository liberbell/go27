package main

import "fmt"

func main() {
	userName := make([]string, 2)

	userName = append(userName, "Max")
	userName = append(userName, "Sam")
	fmt.Println(userName)
}
