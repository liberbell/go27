package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() float64 {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		fmt.Println("Error reading balance file: ", err)
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		fmt.Println("Error parsing balance: ", err)
	}
	return balance
}

func writeBalanceTofile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance = 1000.0
	var depostAmount float64
	var withdrawAmount float64

	fmt.Println("Welcome to Go Bank.")
	fmt.Println("What do you want to do?")

	for {
		fmt.Println("1 Check balance")
		fmt.Println("2 Deposit money")
		fmt.Println("3 Withdraw money")
		fmt.Println("4 Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		// wantCheckBBalance := choice == 1

		switch choice {
		case 1:
			fmt.Println("Your Balance is ", accountBalance)

		case 2:
			fmt.Print("How much do you want to deposit? ")
			fmt.Scan(&depostAmount)
			if depostAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than zero.")
				continue
			}

			accountBalance = accountBalance + depostAmount
			fmt.Printf("Depost %.1f. New Amount is %.1f\n", depostAmount, accountBalance)
			writeBalanceTofile(accountBalance)

		case 3:
			fmt.Print("How much do you want to withdraw? ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than zero.")
				continue
			}

			if withdrawAmount <= accountBalance {
				accountBalance -= withdrawAmount
				fmt.Printf("Withdraw Amount is %.1f. New Amount is %.1f\n", withdrawAmount, accountBalance)
				writeBalanceTofile(accountBalance)
			} else {
				fmt.Println("Ooops! Your are poor.")
				continue
			}
		default:
			fmt.Println("Goodbye!")
			return
		}

		// if choice == 1 {
		// 	fmt.Println("Your Balance is ", accountBalance)
		// } else if choice == 2 {
		// 	fmt.Print("How much do you want to deposit? ")
		// 	fmt.Scan(&depostAmount)
		// 	if depostAmount <= 0 {
		// 		fmt.Println("Invalid amount. Must be greater than zero.")
		// 		continue
		// 	}

		// 	accountBalance = accountBalance + depostAmount
		// 	fmt.Printf("Depost %.1f. New Amount is %.1f\n", depostAmount, accountBalance)
		// } else if choice == 3 {
		// 	fmt.Print("How much do you want to withdraw? ")
		// 	fmt.Scan(&withdrawAmount)
		// 	if withdrawAmount <= 0 {
		// 		fmt.Println("Invalid amount. Must be greater than zero.")
		// 		return
		// 	}

		// 	if withdrawAmount <= accountBalance {
		// 		accountBalance -= withdrawAmount
		// 		fmt.Printf("Withdraw Amount is %.1f. New Amount is %.1f\n", withdrawAmount, accountBalance)
		// 	} else {
		// 		fmt.Println("Ooops! Your are poor.")
		// 	}
		// } else if choice == 4 {
		// 	fmt.Println("Goodbye!")
		// 	return
		// }
		// }
	}
	fmt.Println("Thanks for choosing our bank.")
}
