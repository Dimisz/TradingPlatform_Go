package main

import "fmt"

func printMenu() {
	fmt.Println("1: Print help")
	fmt.Println("2: Print exchange stats")
	fmt.Println("3: Place an ask")
	fmt.Println("4: Place a bid")
	fmt.Println("5: Print wallet")
	fmt.Println("6: Go to the next time frame")
	fmt.Println("=============================")
	fmt.Println("Type in 1 - 6")
}

func printHelp() {
	fmt.Println("Help - Your goal is to make money.")
	fmt.Println("Analyze the market and make bids. Or don't.")
}

func printMarketStats() {
	fmt.Println("The market is horrible.")
}

func makeOffer() {
	fmt.Println("Make an offer - enter the amount")
}

func makeBid() {
	fmt.Println("Make a bid - enter the amount")
}

func printWallet() {
	fmt.Println("Your wallet is empty")
}

func goToNextTimeFrame() {
	fmt.Println("Going to next time frame...")
}
