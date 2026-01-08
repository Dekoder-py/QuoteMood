package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	options := [2]string{"happy", "sad"}

	fmt.Println("Welcome to QuoteMood!")
	fmt.Println("Mood Options:")
	for _, opt := range options {
		fmt.Println("- " + opt)
	}

	fmt.Println("Enter your mood:")
	fmt.Print(">> ")

	// get mood from user input
	reader := bufio.NewReader(os.Stdin)
	mood, _ := reader.ReadString('\n')
	mood = strings.TrimSpace(strings.ToLower(mood))

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
