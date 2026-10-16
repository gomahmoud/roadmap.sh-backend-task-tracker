package main

import (
	"errors"
	"strconv"
)

var availableArgs = map[string]int{
	"add":              1,
	"update":           2,
	"delete":           1,
	"mark-in-progress": 1,
	"mark-done":        1,
	"list":             1,
}

// Helpers

func parseId(arg string) (int, error) {
	id, err := strconv.Atoi(arg)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// APIs

func Validate(commandLine []string) error {
	cmd, args := commandLine[0], commandLine[1:]
	if availableArgs[cmd] == 0 {
		return errors.New("Invalid command: " + cmd)
	}
	if cmd == "list" {
		if len(args) > 1 {
			return errors.New("Too many arguments for command: " + cmd)
		}
		return nil
	}
	if len(args) > availableArgs[cmd] {
		return errors.New("Too many arguments for command: " + cmd)
	}
	if len(args) < availableArgs[cmd] {
		return errors.New("Not enough arguments for command: " + cmd)
	}
	return nil
}

func Parse(commandLine []string) error {
	cmd, args := commandLine[0], commandLine[1:]
	switch cmd {
	case "add":
		return CreateTask(args[0])
	case "update":
		if id, err := parseId(args[0]); err != nil {
			return err
		} else {
			return UpdateTask(id, args[1])
		}
	case "delete":
		if id, err := parseId(args[0]); err != nil {
			return err
		} else {
			return DeleteTask(id)
		}
	case "mark-in-progress":
		if id, err := parseId(args[0]); err != nil {
			return err
		} else {
			return MarkTask(id, TaskInProgress)
		}
	case "mark-done":
		if id, err := parseId(args[0]); err != nil {
			return err
		} else {
			return MarkTask(id, TaskDone)
		}
	case "list":
		if len(args) == 0 {
			return ListTaks("all")
		} else {
			return ListTaks(TaskStatus(args[0]))
		}
	default:
		return errors.New("unknown error")
	}
}
