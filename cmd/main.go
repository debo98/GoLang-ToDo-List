package main

import (
	"bufio"
	"fmt"
	"os"
)

func fallbackMessage(_ *[]task, reader *bufio.Reader) {
	fmt.Printf("Please enter a digit between 1-4\n")
}

func performFunction(choice string) func(*[]task, *bufio.Reader) {
	switch choice {
		case "1": 
			return addTask
		case "2":
			return viewTasks
		case "3":
			return completeTask
		case "4":
			return exitApp
		}
	return fallbackMessage
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	tasks := getTasks()

	for {
		fmt.Print(mainMenu)
		userChoice := takeSingleLineInput(reader)

		var functionToExecute = performFunction(userChoice)
		functionToExecute(&tasks, reader)
		fmt.Println("-----------------------------")
	}
}
