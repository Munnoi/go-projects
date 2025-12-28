package main

import (
	"fmt"
	"os"
	"encoding/json"
	"bufio"
	"strings"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

func addTask() {
	// Reading the task input.
	reader := bufio.NewReader(os.Stdin)
	var taskName string
	fmt.Print("Enter task name: ")
	taskName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	taskName = strings.TrimSpace(taskName) // Trim whitespaces.

	// Reading the tasks.
	file, err := os.ReadFile("db.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var tasks TaskList
	err = json.Unmarshal(file, &tasks) // Converting json to TaskList format.
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Assigning an ID to the new task.
	nextID := 1
	if len(tasks.Tasks) > 0 {
		nextID = tasks.Tasks[len(tasks.Tasks)-1].ID + 1 // ID+1 of the last task ID.
	}

	// Creating the new task.
	newTask := Task{ID: nextID, Name: taskName, Done: false}
	tasks.Tasks = append(tasks.Tasks, newTask)

	// Writing the updated tasks to the file.
	updatedData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Writing the updated data to the file.
	err = os.WriteFile("db.json", updatedData, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Task added successfully!")
}

func viewTasks() {
	file, err := os.Open("db.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read and print tasks
	// For simplicity, let's assume each task is on a new line
	// In a real application, you might want to parse JSON
	buffer := make([]byte, 1024)
	n, err := file.Read(buffer)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var tasks TaskList	

	err = json.Unmarshal(buffer[:n], &tasks)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	// Display tasks
	for _, task := range tasks.Tasks {
		status := "❌"
		if task.Done {
			status = "✅"
		}
		fmt.Printf("%d: %s %s\n", task.ID, task.Name, status)
	}
}

func completeTask() {
	fmt.Println("Complete task functionality not yet implemented.")
}

func deleteTask() {
	fmt.Println("Delete task functionality not yet implemented.")
}

func displayMenu() {
	fmt.Println("1. Add task.")
	fmt.Println("2. View tasks.")
	fmt.Println("3. Complete task.")
	fmt.Println("4. Delete task.")
	fmt.Println("5. Exit.")
	fmt.Print("Enter your choice: ")
	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		addTask()
	case 2:
		fmt.Println("\nTasks:")
		viewTasks()
		fmt.Println()
	case 3:
		completeTask()
	case 4:
		deleteTask()
	case 5:
		os.Exit(0)
	default:
		fmt.Println("Invalid choice. Please try again.")
	}
}

func TaskManager() {
	for true {
		displayMenu()
	}
}

func main() {
	TaskManager()
}
