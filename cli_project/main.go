package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

// Task represents a single to-do item
type Task struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

const taskFile = "tasks.json"

func main() {
	// 1. DEFINE FLAGS
	add := flag.String("add", "", "Add a new task")
	list := flag.Bool("list", false, "List all tasks")
	deleteID := flag.Int("delete", 0, "Delete task by ID")
	flag.Parse()

	// 2. LOAD EXISTING TASKS
	tasks, err := loadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		os.Exit(1)
	}

	// 3. HANDLE COMMANDS
	switch {
	case *add != "":
		newTask := Task{
			ID:   len(tasks) + 1,
			Text: *add,
			Done: false,
		}
		tasks = append(tasks, newTask)
		saveTasks(tasks)
		fmt.Printf("✅ Added: %s\n", *add)

	case *list:
		fmt.Println("📋 Your Task List:")
		if len(tasks) == 0 {
			fmt.Println("   (No tasks found. Try adding one!)")
			return
		}
		for _, t := range tasks {
			status := " "
			if t.Done {
				status = "x"
			}
			fmt.Printf("   [%s] %d: %s\n", status, t.ID, t.Text)
		}

	case *deleteID != 0:
		newTasks := []Task{}
		found := false
		for _, t := range tasks {
			if t.ID == *deleteID {
				found = true
				continue
			}
			newTasks = append(newTasks, t)
		}
		if found {
			saveTasks(newTasks)
			fmt.Printf("🗑️ Deleted task with ID: %d\n", *deleteID)
		} else {
			fmt.Printf("❌ Task with ID %d not found.\n", *deleteID)
		}

	default:
		flag.Usage()
	}
}

// Helper to load tasks from JSON file
func loadTasks() ([]Task, error) {
	file, err := os.ReadFile(taskFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}

	var tasks []Task
	if err := json.Unmarshal(file, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

// Helper to save tasks to JSON file
func saveTasks(tasks []Task) {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Printf("Error encoding tasks: %v\n", err)
		return
	}
	if err := os.WriteFile(taskFile, data, 0644); err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
	}
}

/*
SUMMARY:
- We used the 'flag' package to create a real CLI interface.
- We handled persistent storage using 'os' and 'encoding/json'.
- We built a simple CRUD-like logic (Create, Read, Delete).
- This project integrates everything we've learned: variables, slices, maps, structs, and control flow!
*/
