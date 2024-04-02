package main

import "fmt"

func main() {
	var accountBalance = 1000.0
	var depostBalance float64

	fmt.Println("Welcome to Go Bank.")
	fmt.Println("What do you want to do?")
	fmt.Println("1 Check balance")
	fmt.Println("2 Deposit money")
	fmt.Println("3 Withdraw money")
	fmt.Println("4 Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	// wantCheckBBalance := choice == 1

	if choice == 1 {
		fmt.Println("Your Balance is ", accountBalance)
	} else if choice == 2 {
		fmt.Print("How much do you want to deposit?")
		fmt.Scan(&depostBalance)
		if accountBalance > depostBalance {
			accountBalance = accountBalance - depostBalance
			fmt.Printf("Depost %.1f. Your Balance is %.1f\n", depostBalance, accountBalance)
		} else {
			fmt.Println("Your Balance is too low.")
		}
	}

	fmt.Println("Your choice is: ", choice)
}
