package main

import "fmt"

func main() {
	var options = [2]string{"happy", "sad"}

	fmt.Println("Welcome to QuoteMood!")
	fmt.Println("Mood Options:")
	for _, opt := range options{
	fmt.Println("- " + opt)	
	}

	fmt.Println("Enter your mood:")
	fmt.Print(">> ")

	// get mood from user input
	var mood string
	fmt.Scan(&mood)

	// process mood - must be one of the options else ask again.
	switch mood {
	case "happy":
		fmt.Println("Happy quote")	
	case "sad":
		fmt.Println("Sad quote")
	default:
		fmt.Println("Sorry, we didn't recognise that answer. Please choose one of the options.")
	}
}
