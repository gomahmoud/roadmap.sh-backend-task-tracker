# CLI Task Tracker

A small command-line application for managing tasks locally. Tasks are stored in a JSON file, so they remain available between runs without requiring a database or network connection.

## Requirements

- Go 1.27.1 or later

## Run the application

From the project directory, run commands with `go run .`:

```sh
go run . add "Buy groceries"
go run . list
```

To build an executable:

```sh
go build -o task-cli .
```

Then run it using the same commands, replacing `go run .` with `./task-cli` (on Windows, use `task-cli.exe`).

Run the application from the directory where you want `tasks.json` to be stored. The data file path is relative to the current working directory.

## Commands

| Command | Description | Example |
| --- | --- | --- |
| `add <description>` | Create a task with `todo` status | `go run . add "Buy groceries"` |
| `update <id> <description>` | Change a task's description | `go run . update 1 "Buy groceries and cook dinner"` |
| `delete <id>` | Delete a task | `go run . delete 1` |
| `mark-in-progress <id>` | Mark a task as in progress | `go run . mark-in-progress 1` |
| `mark-done <id>` | Mark a task as done | `go run . mark-done 1` |
| `list` | Show all tasks | `go run . list` |
| `list <status>` | Show tasks with the given status | `go run . list todo` |

The supported statuses are `todo`, `in-progress`, and `done`. Listing can be filtered with any of these values. Task IDs are assigned using the highest existing ID plus one.

Descriptions containing spaces should be enclosed in quotes.

## Data storage

The application stores tasks in `tasks.json` as a JSON array. If the file does not exist, it is initialized as an empty array.
Each task has an ID, description, status, creation timestamp, and update timestamp. Timestamps use RFC 3339 format under the `created_at` and `updated_at` JSON fields.

Updates are written through a temporary `tasks.tmp.json` file and then moved into place as `tasks.json`.

## Errors

Invalid commands, incorrect argument counts, non-numeric IDs, missing tasks when updating or marking them, and file or JSON errors are reported on the command line.
