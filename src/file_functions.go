package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func takeSingleLineInput(reader *bufio.Reader) string {
	inputVar, _ := reader.ReadString('\n')
	inputVar = strings.TrimSuffix(inputVar, "\n")
	return inputVar
}

func loadTasksStringFromFile() string {
	tasks := ""
	
	// Read the file
	tasksString, err := os.ReadFile(tasksFileName)
	if err != nil {
		if os.IsNotExist(err) {
			// If file doesn't exist, return empty task list
			return tasks
		}
		fmt.Printf("Error reading tasks file: %s\n", err.Error())
		return tasks
	}
	return string(tasksString)
}

func parseTasksString(tasksString string) []task {
	tasks := []task{}

	// If tasksString is empty, return an empty slice of tasks
	if tasksString == "" { return tasks }

	// Split tasksString into lines and add non-empty lines to tasks
	lines := strings.Split(tasksString, "\n")
	for _ , line := range lines {
		if line != "" {
			csvParsedLine := strings.Split(string(line), ";")
			index, _ := strconv.Atoi(csvParsedLine[0])  // Will not give error as I am saving this using this file itself, assume nobody else touches it
			taskInLine := task{uint8(index), csvParsedLine[1], bool(csvParsedLine[2] == "true")}
			tasks = append(tasks, taskInLine)
		}
	}
	return tasks
}

func getTasks() []task {
	tasksString := loadTasksStringFromFile()
	tasks := parseTasksString(tasksString)
	return tasks
}

func getTasksAsASingleString(tasks *[]task) string {
	var stringifiedTasks []string
	// Create a slice of file content, each element containing the details of a single task
	// Save format: "index;name;isCompleted" -> semi-colon separated values
	for _, task := range *tasks{
		stringifiedTask := strconv.Itoa(int(task.index)) + ";" + task.name + ";" + strconv.FormatBool(task.isCompleted)
		stringifiedTasks = append(stringifiedTasks, stringifiedTask)
	}

	// Make a single line out of this slice to save to file
	tasksAsASingleString := strings.Join(stringifiedTasks, "\n")
	return tasksAsASingleString
}

func saveTasksToFile(tasksAsASingleString string) {
	err := os.WriteFile(tasksFileName, []byte(tasksAsASingleString), 0600)  // Read-Write permissions for user, none otherwise
	if err != nil {
		fmt.Printf("Error saving tasks to file: %s\n", err.Error())
	}
}

func saveTasks(tasks *[]task) {
	tasksAsASingleString := getTasksAsASingleString(tasks)
	saveTasksToFile(tasksAsASingleString)
	
}
