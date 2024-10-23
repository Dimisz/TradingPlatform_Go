package main

import (
	"fmt"
	"strconv"
)

func getUserOption() int {
	var optionChosen int
	var userInput string
	for {
		_, err := fmt.Scanln(&userInput)
		if err != nil {
			fmt.Println("invalid input")
			continue
		}
		optionChosen, err = strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("could not parse input")
			continue
		}
		if optionChosen >= 1 && optionChosen <= 6 {
			return optionChosen
		}
		fmt.Println("Choose a valid option from 1 to 6")
		printMenu()
	}
}

func processUserOption() {
	option := getUserOption()
	fmt.Printf("You chose %d:\n", option)
	switch option {
	case 1:
		printHelp()
	case 2:
		printMarketStats()
	case 3:
		makeOffer()
	case 4:
		makeBid()
	case 5:
		printWallet()
	case 6:
		goToNextTimeFrame()
	default:
		fmt.Println("Bad input. Try again")
	}
}
