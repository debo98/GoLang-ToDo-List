# Go To-Do List Application

This is a simple command-line To-Do List application written in Go. It allows users to manage their tasks by adding, viewing, and marking them as completed.

## Features

- **Add a new task**: Users can input a new task to be added to their list.
- **View all tasks**: Users can see all the tasks they have added, along with their completion status.
- **Mark a task as completed**: Users can mark a specific task as completed by entering its index.
- **Exit the application**: Users can exit the application gracefully.
- **File Handling**: The tasks are saved into a text file `tasks.txt` when the application is exited gracefully. On start, the tasks are loaded from this file. Each task is stored on a new line, with its completion status indicated.

## Getting Started

### Prerequisites

- Go (version 1.16 or later) installed on your machine.

### Installation

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd <repository-directory>
   ```

2. Build the application:

   ```bash
   go build main.go
   ```

### Usage

1. Run the application:

   ```bash
   ./main
   ```

2. Follow the on-screen menu to manage your tasks:
   - Enter `1` to add a new task.
   - Enter `2` to view all tasks.
   - Enter `3` to mark a task as completed.
   - Enter `4` to exit the application.

