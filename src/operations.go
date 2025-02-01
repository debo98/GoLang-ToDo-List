package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func addTask(tasks *[]task, reader *bufio.Reader) {
	// Ask user for the task to be added
	fmt.Println("What is the task that you want to add?")

	newTask := takeSingleLineInput(reader)

	// Add this new task to the existing slice of tasks
	newTaskIndex := uint8(len(*tasks))  // Since this is 0-based indexing
	*tasks = append(*tasks, task{newTaskIndex, newTask, false})
}

func viewTasks(tasks *[]task, reader *bufio.Reader) {
	fmt.Printf("This is the list of tasks:\n")

	const checklistSymbolIsCompleted string = "[Done]"
	const checklistSymbolIsNotCompleted string = "[    ]"
	// Iterate through all the tasks and print them
	for _, task := range *tasks{
		var symbol string
		if task.isCompleted {
			symbol = checklistSymbolIsCompleted
		} else {
			symbol = checklistSymbolIsNotCompleted
		}
		fmt.Printf("%s %s\n", symbol, task.name)
	}
}

func completeTask(tasks *[]task, reader *bufio.Reader) {
	fmt.Printf("Which task have you completed? Enter the index of the task\n")

	taskIndexString := takeSingleLineInput(reader)

	taskIndex, err := strconv.Atoi(taskIndexString)
	if err != nil{
		fmt.Println("Please enter the index corresponding to the task. Doing Nothing!")
	}
	if taskIndex > len(*tasks) || taskIndex <= 0 {
		fmt.Println("This is an invalid index. Doing Nothing!")
	} else if (*tasks)[taskIndex-1].isCompleted {
		fmt.Printf("The task '%s' has already been marked as done\n", (*tasks)[taskIndex-1].name)
	} else {
		(*tasks)[taskIndex-1].isCompleted = true
	}
}

func exitApp(tasks *[]task, reader *bufio.Reader) {
	saveTasks(tasks)
	os.Exit(0)
}