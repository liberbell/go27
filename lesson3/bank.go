package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading balance file: ", err)
		return 1000, errors.New("Failed to find file.")
	}
	valueText := string(data)
	balance, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		fmt.Println("Error parsing balance: ", err)
		return 1000, errors.New("Failed to parse stored value")
	}
	return balance, nil
}

func writeFloatTofile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func main() {
	var accountBalance, err = getFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	var depostAmount float64
	var withdrawAmount float64

	fmt.Println("Welcome to Go Bank.")

	for {
		presentOptions()

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
			writeFloatTofile(accountBalance, accountBalanceFile)

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
				writeFloatTofile(accountBalance, accountBalanceFile)
			} else {
				fmt.Println("Ooops! Your are poor.")
				continue
			}
		default:
			fmt.Println("Goodbye!")
			return
		}
	}
	fmt.Println("Thanks for choosing our bank.")
}
