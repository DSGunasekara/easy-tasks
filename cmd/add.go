package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var tasksFile = "tasks.json"

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [task name]",
	Short: "Add a new task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskName := args[0]
		addTask(taskName)
	},
}

func addTask(name string) {
	tasks := loadTasks()

	task := Task{ID: len(tasks) + 1, Name: name}
	tasks = append(tasks, task)

	saveTasks(tasks)
	fmt.Printf("✅ Task added: %s\n", name)
}

func loadTasks() []Task {
	data, err := os.ReadFile(tasksFile)
	if err != nil {
		return []Task{}
	}

	var tasks []Task
	_ = json.Unmarshal(data, &tasks)
	return tasks
}

func saveTasks(tasks []Task) {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	_ = os.WriteFile(tasksFile, data, 0644)
}

func init() {
	rootCmd.AddCommand(addCmd)
}
