package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	options := [8]string{"happy", "sad", "angry", "mad", "tired", "excited", "annoyed", "hopeful"}

	fmt.Println("Welcome to QuoteMood!")

	for true { // run forever - processMood exits the program
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
		for _, opt := range options {
			if mood == opt {
				processMood(mood)
			}
		}
		fmt.Println("Sorry, we didn't recognise that answer. Please choose one of the options.")
	}
}

func processMood(mood string) {
	fmt.Println(mood + " quote")
	os.Exit(0)
}
