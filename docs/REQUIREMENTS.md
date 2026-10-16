# Task Tracker CLI — Requirements

## 1. Project Overview

A command-line application for managing tasks locally.

## 2. Scope

- Add, update, and delete tasks.
- Change task status.
- List tasks by status.
- Store tasks in a JSON file.

## 3. Functional Requirements

- FR-01: Add a task.
- FR-02: Update a task description.
- FR-03: Permanently delete a task.
- FR-04: Mark a task as in-progress.
- FR-05: Mark a task as done.
- FR-06: List all tasks.
- FR-07: Filter tasks by status.
- FR-08: Filter tasks that are not done.

## 4. Data Requirements

- id
- description
- status
- createdAt
- updatedAt

## 5. Storage Rules

- Filename: tasks.json
- Format: JSON array of task objects.
- Create the file if it does not exist.
- Generate IDs using the highest existing ID + 1.
- Never overwrite malformed JSON with an empty task list.

## 6. Input Validation and Errors

- Reject empty or whitespace-only descriptions.
- Handle invalid commands and arguments.
- Handle nonexistent task IDs.
- Report file-reading and file-writing errors.

## 7. Acceptance Criteria

Define how to verify that every requirement works.
