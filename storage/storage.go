package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"TaskCLI/task"
)

const filename = "tasks.json"

type SaveData struct {
	Tasks  []task.Task `json:"tasks"`
	NextID int         `json:"next_id"`
}

func SaveTasks(tl *task.TaskList) error {
	data := SaveData{
		Tasks:  tl.Tasks,
		NextID: tl.NextID,
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to convert tasks to JSON: %w", err)
	}

	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func LoadTasks() (*task.TaskList, error) {
	jsonData, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return task.NewTaskList(), nil
		}
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var data SaveData
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	taskList := &task.TaskList{
		Tasks:  data.Tasks,
		NextID: data.NextID,
	}

	return taskList, nil
}