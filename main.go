package main

import (
	"fmt"
	"time"
)

type Task struct {
	summary     string
	description string
	deadline    time.Time
	priority    int
}

type Note struct {
	title string
	text  string
}

type ToDoList struct {
	name  string
	tasks []Task
	notes []Note
}

func (t Task) IsOverdue() bool {
	return time.Now().After(t.deadline)
}

func (t Task) IsTopPriority() bool {
	return t.priority >= 4
}

func (l ToDoList) TasksCount() int {
	return len(l.tasks)
}

func (l ToDoList) NotesCount() int {
	return len(l.notes)
}

func (l ToDoList) CountTopPrioritiesTasks() int {
	var result int
	for _, i := range l.tasks {
		if i.priority >= 4 {
			result++
		}
	}

	return result
}

func (l ToDoList) CountOverdueTasks() int {
	var result int
	for _, i := range l.tasks {
		if i.IsOverdue() {
			result++
		}
	}

	return result
}

func main() {
	task := Task{summary: "Make Yandex Lyceum", deadline: time.Now().Add(time.Hour), description: "Make Module 0, Task 10", priority: 4}
	fmt.Println(task.IsOverdue())
	fmt.Print(task.IsTopPriority())
}
