package main

import (
	"fmt"
	"strconv"
	"os"
	"TaskCLI/task"
	"TaskCLI/storage"
)


func main() {
	taskList, err := storage.LoadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "add":
		handleAdd(taskList, os.Args[2:])
	case "list":
		handleList(taskList)
	case "complete":
		handleComplete(taskList, os.Args[2:])
	case "delete":
		handleDelete(taskList, os.Args[2:])
	case "update":
		handleUpdate(taskList, os.Args[2:])
	case "help":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n\n", command)
		printUsage()
		os.Exit(1)
	}
}

func handleAdd(tl *task.TaskList, args []string) {
	if len(args) == 0 {
		fmt.Println("Error: Please provide a task description")
		fmt.Println("Usage: TaskCLI add \"Your task description\"")
		os.Exit(1)
	}

	description := args[0]
	for i := 1; i < len(args); i++ {
		description += " " + args[i]
	}

	newTask := tl.AddTask(description)

	err := storage.SaveTasks(tl)
	if err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Task added successfully (ID: %d)\n", newTask.ID)
}

func handleList(tl *task.TaskList) {
	tasks := tl.ListTasks()

	if len(tasks) == 0 {
		fmt.Println("No tasks found. Add one with: TaskCLI add \"Your task\"")
		return
	}

	fmt.Println("\n Your Tasks:")
	fmt.Println("─────────────────────────────────────────────────────")
	for _, t := range tasks {
		statusIcon := getStatusIcon(t.Status)
		fmt.Printf("[%d] %s %s - %s\n", t.ID, statusIcon, t.Description, t.Status)
	}
	fmt.Println("─────────────────────────────────────────────────────")

	todoCount := 0
	inProgressCount := 0
	doneCount := 0

	for _, t := range tasks {
		switch t.Status {
		case "todo":
			todoCount++
		case "in-progress":
			inProgressCount++
		case "done":
			doneCount++
		}
	}
	
	fmt.Printf("Total: %d | Todo: %d | In Progress: %d | Done: %d\n\n", 
		len(tasks), todoCount, inProgressCount, doneCount)
}

func handleComplete(tl *task.TaskList, args []string) {
	if len(args) == 0 {
		fmt.Println("Error: Please provide a task ID")
		fmt.Println("Usage: TaskCLI complete <id>")
		os.Exit(1)
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Error: Invalid task ID '%s'. Please provide a number.\n", args[0])
		os.Exit(1)
	}

	err = tl.CompleteTask(id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	err = storage.SaveTasks(tl)
	if err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Task %d marked as complete!\n", id)
}

func handleDelete(tl *task.TaskList, args []string) {
	if len(args) == 0 {
		fmt.Println("Error: Please provide a task ID")
		fmt.Println("Usage: TaskCLI delete <id>")
		os.Exit(1)
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Error: Invalid task ID '%s'. Please provide a number.\n", args[0])
		os.Exit(1)
	}

	err = tl.DeleteTask(id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	err = storage.SaveTasks(tl)
	if err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Task %d deleted successfully!\n", id)
}

func handleUpdate(tl *task.TaskList, args []string) {
	if len(args) < 2 {
		fmt.Println("Error: Please provide a task ID and new description")
		fmt.Println("Usage: TaskCLI update <id> \"New description\"")
		os.Exit(1)
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Error: Invalid task ID '%s'. Please provide a number.\n", args[0])
		os.Exit(1)
	}

	newDescription := args[1]
	for i := 2; i < len(args); i++ {
		newDescription += " " + args[i]
	}

	err = tl.UpdateTask(id, newDescription)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	err = storage.SaveTasks(tl)
	if err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Task %d updated successfully!\n", id)
}

func getStatusIcon(status string) string {
	switch status {
	case "todo":
		return "[ ]"
	case "in-progress":
		return "[→]"
	case "done":
		return "[✓]"
	default:
		return "[ ]"
	}
}

func printUsage() {
	fmt.Println("Task CLI - A simple CLI to-do list manager")
	fmt.Println("\nUsage:")
	fmt.Println("  TaskCLI <command> [arguments]")
	fmt.Println("\nAvailable Commands:")
	fmt.Println("  add <description>       Add a new task")
	fmt.Println("  list                    List all tasks")
	fmt.Println("  complete <id>           Mark a task as complete")
	fmt.Println("  delete <id>             Delete a task")
	fmt.Println("  update <id> <desc>      Update task description")
	fmt.Println("  help                    Show this help message")
	fmt.Println("\nExamples:")
	fmt.Println("  TaskCLI add \"Buy groceries\"")
	fmt.Println("  TaskCLI list")
	fmt.Println("  TaskCLI complete 1")
	fmt.Println("  TaskCLI update 2 \"Buy milk and eggs\"")
	fmt.Println("  TaskCLI delete 3")
	fmt.Println()
}