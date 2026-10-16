package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	dbPath    = "./tasks.json"
	tmpDBPath = "./tasks.tmp.json"
)

// helpers

func readTasks() ([]Task, error) {
	file, err := os.Open(dbPath)
	if err != nil {
		return nil, fmt.Errorf("error opening tasks file: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("error closing tasks file: %v\n", err)
		}
	}(file)

	var tasks []Task
	if err := json.NewDecoder(file).Decode(&tasks); err != nil {
		return nil, fmt.Errorf("error decoding tasks JSON: %w", err)
	}
	return tasks, nil
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		return fmt.Errorf("error encoding tasks JSON: %w", err)
	}

	tmpFile, err := os.Create(tmpDBPath)
	if err != nil {
		return fmt.Errorf("error creating temporary tasks file: %w", err)
	}

	tmpPath := tmpFile.Name()
	defer func() {
		_ = os.Remove(tmpPath)
	}()

	if _, err := tmpFile.Write(data); err != nil {
		_ = tmpFile.Close()
		return fmt.Errorf("error writing temporary tasks file: %w", err)
	}

	if err := tmpFile.Sync(); err != nil {
		_ = tmpFile.Close()
		return fmt.Errorf("error syncing temporary tasks file: %w", err)
	}

	if err := tmpFile.Close(); err != nil {
		return fmt.Errorf("error closing temporary tasks file: %w", err)
	}

	if err := os.Rename(tmpPath, dbPath); err != nil {
		return fmt.Errorf("error replacing tasks file: %w", err)
	}

	return nil
}

func filterTasks(tasks []Task, filter TaskStatus) []Task {
	var filtered []Task
	for _, task := range tasks {
		if task.Status == filter {
			filtered = append(filtered, task)
		}
	}
	return filtered
}

// APIs

func InsertOne(task Task) error {
	tasks, err := readTasks()
	if err != nil {
		return fmt.Errorf("error reading tasks: %w", err)
	}
	tasks = append(tasks, task)

	if err := saveTasks(tasks); err != nil {
		return fmt.Errorf("error saving tasks: %w", err)
	}
	return nil
}

func UpdateOne(task Task) error {
	tasks, err := readTasks()
	if err != nil {
		return fmt.Errorf("error reading tasks: %w", err)
	}
	for i, t := range tasks {
		if t.Id == task.Id {
			tasks[i] = task
			break
		}
	}
	if err := saveTasks(tasks); err != nil {
		return fmt.Errorf("error saving tasks: %w", err)
	}
	return nil
}

func DeleteOne(id int) error {
	tasks, err := readTasks()
	if err != nil {
		return fmt.Errorf("error reading tasks: %w", err)
	}
	for i, task := range tasks {
		if task.Id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
	if err := saveTasks(tasks); err != nil {
		return fmt.Errorf("error saving tasks: %w", err)
	}
	return nil
}

func GetOneById(id int) (Task, error) {
	tasks, err := readTasks()
	if err != nil {
		return Task{}, fmt.Errorf("error reading tasks: %w", err)
	}
	for _, task := range tasks {
		if task.Id == id {
			return task, nil
		}
	}
	return Task{}, fmt.Errorf("task not found")
}

func GetHighestId() (int, error) {
	tasks, err := readTasks()
	if err != nil {
		return 0, fmt.Errorf("error reading tasks: %w", err)
	}
	highestId := 0
	for _, t := range tasks {
		if t.Id > highestId {
			highestId = t.Id
		}
	}
	return highestId, nil
}

func GetAll(filter TaskStatus) ([]Task, error) {
	tasks, err := readTasks()
	if err != nil {
		return nil, fmt.Errorf("error reading tasks: %w", err)
	}
	if filter == "all" {
		return tasks, nil
	}
	tasks = filterTasks(tasks, filter)
	return tasks, nil
}
