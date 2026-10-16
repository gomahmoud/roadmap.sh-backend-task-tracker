package main

import (
	"fmt"
	"time"
)

// Helpers

func printTask(task Task) {
	fmt.Println("id:          ", task.Id)
	fmt.Println("description: ", task.Description)
	fmt.Println("status:      ", task.Status)
	fmt.Println("createdAt:   ", task.CreatedAt)
	fmt.Println("updatedAt:   ", task.UpdatedAt)
	fmt.Println()
}

// APIs

func CreateTask(description string) error {
	highestId, err := GetHighestId()
	if err != nil {
		return fmt.Errorf("Error getting highest id: %v", err)
	}
	var task = Task{
		Id:          highestId + 1,
		Description: description,
		Status:      TaskTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}

	if err := InsertOne(task); err != nil {
		return fmt.Errorf("Error inserting task: %v", err)
	}
	return nil
}

func UpdateTask(id int, description string) error {
	task, err := GetOneById(id)
	if err != nil {
		return fmt.Errorf("Error getting task: %v", err)
	}
	task.Description = description
	task.UpdatedAt = time.Now()
	if err := UpdateOne(task); err != nil {
		return fmt.Errorf("Error updating task: %v", err)
	}
	return nil
}

func DeleteTask(id int) error {
	err := DeleteOne(id)
	if err != nil {
		return fmt.Errorf("Error deleting task: %v", err)
	}
	return nil
}

func MarkTask(id int, status TaskStatus) error {
	task, err := GetOneById(id)
	if err != nil {
		return fmt.Errorf("Error getting task: %v", err)
	}
	task.Status = status
	task.UpdatedAt = time.Now()
	if err := UpdateOne(task); err != nil {
		return fmt.Errorf("Error updating task: %v", err)
	}
	return nil
}

func ListTaks(filter TaskStatus) error {
	tasks, err := GetAll(filter)
	if err != nil {
		return fmt.Errorf("Error getting tasks: %v", err)
	}
	for _, task := range tasks {
		printTask(task)
	}
	return nil
}
