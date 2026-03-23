# CLI Task Manager

A simple, fast, and local task manager built with Go and Cobra.
Tasks are saved as a JSON file in your home directory `~/.tasks.json`.

## Installation

```bash
go build -o cli-task
```

## Usage

### 1. Add a task
```bash
./cli-task add Clean the house
./cli-task add Buy groceries
./cli-task add Walk the dog
```

### 2. List tasks
```bash
./cli-task list
```

### 3. Complete a task
```bash
./cli-task do 1
```
(where 1 is the Task ID shown in `list`)

### 4. Delete a task
```bash
./cli-task delete 2
```

### Project URL https://github.com/kris5885/cli-task
