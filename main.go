package main

import "fmt"

func main() {
	for {
		fmt.Println("pokedex >")

		var userInput string

		_, err := fmt.Scanln(&userInput)
		if err != nil {
			fmt.Println("That's not a valid command")
		}

		if userInput == "help" {
			fmt.Println("This message describes how to use the podedex")
		}
		if userInput == "exit" {
			fmt.Println("Exiting program")
			return
		}
	}
}
