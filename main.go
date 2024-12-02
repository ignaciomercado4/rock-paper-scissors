package main

import (
	"fmt"
	"math/rand"
)

func getMachineChoice(options [3]string) string {
	randomInt := rand.Intn(3)

	return options[randomInt]
}

func getUserChoice (options [3]string) string {
	var userChoice int
	var finalUserChoice string

	fmt.Println("Select an option: \n+ 0 for Rock \n+ 1 for Paper \n+ 2 for Scissors")
	fmt.Scan(&userChoice)

	for i:= range 3 {
		if userChoice == i {

			finalUserChoice = options[userChoice]
			}
		}

	return finalUserChoice
}

func getGameResult(userChoice string, machineChoice string) string {
	var result string = "UNKNOWN"
	
	if userChoice != "" {
		switch userChoice {
		case "Rock":
			if machineChoice == "Rock" {
				result = "TIE"
			}

			if machineChoice == "Paper" {
				result = "LOSE"
			} else {
				result = "WIN"
			}

		case "Paper":
			if machineChoice == "Rock" {
				result = "WIN"
			}

			if machineChoice == "Paper" {
				result = "TIE"
			} else {
				result = "LOSE"
			}

		case "Scissors": 
			if machineChoice == "Rock" {
				result = "LOSE"
			}

			if machineChoice == "PAPER" {
				result = "WIN"
			} else {
				result = "TIE"
			}

		default:
			result = "UNKNOWN"
		}
	}

	return result
}

func main() {
	options := [3]string{"Rock", "Paper", "Scissors"}
	
	userChoice := getUserChoice(options)
	machineChoice := getMachineChoice(options)
	result := getGameResult(userChoice, machineChoice)

	if result == "UNKNOWN" {
		fmt.Println("INVALID INPUT")
	} else {
		fmt.Println("MACHINE CHOSE", machineChoice, "\n YOU", result)
	}
	
}