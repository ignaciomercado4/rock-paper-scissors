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

	fmt.Println("Select an option: 0 for ROCK - 1 for PAPER - 2 for SCISSORS")
	fmt.Scan(&userChoice)

	for i:= range 3 {
		if userChoice == i {

			finalUserChoice = options[userChoice]
			}
		}

	return finalUserChoice
}

func main() {
	options := [3]string{"Rock", "Paper", "Scissors"}
	fmt.Println(getUserChoice(options))
}