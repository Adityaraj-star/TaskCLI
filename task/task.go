package task

import (
	"fmt"
	"time"
)

type Task struct {
	ID			int 		`json:"id"`
	Description	string		`json:"description"`
	Status		string		`json:"status"`			// "todo", "in-progress", "done"
	CreatedAt	time.Time	`json:"created_at"`
}

type TaskList struct {
	Tasks 	[]Task `json:"tasks"`
	NextID 	int
}

func NewTaskList() *TaskList {
	return &TaskList{
		Tasks: 	[]Task{},
		NextID: 1,
	}
}

func (tl *TaskList) AddTask(description string) Task {
	task := Task{
		ID:				tl.NextID,
		Description: 	description,
		Status: 		"todo",
		CreatedAt: 		time.Now(),			
	}

	tl.Tasks = append(tl.Tasks, task)
	tl.NextID++
	return task
}

func (tl *TaskList) ListTasks() []Task {
	return tl.Tasks
}

func (tl *TaskList) CompleteTask(id int) error {
	for i := range tl.Tasks {
		if tl.Tasks[i].ID == id {
			tl.Tasks[i].Status = "done"
			return nil
		}
	}

	return fmt.Errorf("Task with ID %d not found", id)
}

func (tl *TaskList) DeleteTask(id int) error {
	for i := range tl.Tasks {
		if tl.Tasks[i].ID == id {
			tl.Tasks = append(tl.Tasks[:i], tl.Tasks[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Task with ID %d not found", id)
}

func (tl *TaskList) GetTaskByID(id int) (*Task, error) {
	for i := range tl.Tasks {
		if tl.Tasks[i].ID == id {
			return &tl.Tasks[i], nil
		}
	}

	return nil, fmt.Errorf("Task with ID %d not found", id)
}

func (tl *TaskList) UpdateTask(id int, newDescription string) error {
	for i := range tl.Tasks {
		if tl.Tasks[i].ID == id {
			tl.Tasks[i].Description = newDescription
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}