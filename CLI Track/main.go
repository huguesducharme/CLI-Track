package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreateAt    string `json:"create_at"`
	UpdateAt    string `json:"update_at"`
}

func main() {
	fmt.Println("Task Tracker")
	fmt.Println("Enter a name for tracker")

	var nameTracker string
	fmt.Scan(&nameTracker)
	FileName := nameTracker + ".json"

	file, err := os.OpenFile(FileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	// Checking for an empty file
	if info.Size() == 0 {
		_, err := file.Write([]byte("[]"))
		if err != nil {
			log.Fatal(err)
		}
	}

	for { // Cycle for choose action
		fmt.Println("Choose action: 1.Create task | 2.Delete task | 3.Show all task | 4. Show done task | 5. Show todo task | 6. Show inde-degress task")

		var action int
		fmt.Scan(&action)

		switch action {
		case 1:
			CreateTask(FileName)
			return
		case 2:
			DeleteTask(FileName)
			return
		case 3:
			ShowTask(FileName)
			return
		case 4:
			ShowDoneTask(FileName)
			return
		case 5:
			ShowTodoTask(FileName)
			return
		case 6:
			ShowIdneDegressTask(FileName)
			return
		default:
			fmt.Println("Inccorect action, repeat again...")
		}
	}
}

func CreateTask(FileName string) {
	fmt.Println("Enter a description of the task")
	var description string
	fmt.Scan(&description)

	fmt.Println("Enter task status")
	var status string
	fmt.Scan(&status)

	id := rand.Intn(1000)
	createAt := time.Now().Format("02.01.2006 15:04:05")
	updateAt := time.Now().Format("02.01.2006 15:04:05")

	userTask := Task{ID: id, Description: description, Status: status, CreateAt: createAt, UpdateAt: updateAt}

	// Read tasks from JSON file
	data, err := os.ReadFile(FileName)
	if err != nil {
		log.Fatal(err)
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		log.Fatal(err)
	}

	// Add a new task
	tasks = append(tasks, userTask)

	// Save []Task back to JSON file
	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(FileName, jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteTask(FileName string) {
	// Read tasks from JSON file
	data, err := os.ReadFile(FileName)
	if err != nil {
		log.Fatal(err)
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Enter the task id to delete")
	var targetTask int
	fmt.Scan(&targetTask)

	// Find a task
	var indexTaskToDelete = -1
	for index, task := range tasks {
		if task.ID == targetTask {
			indexTaskToDelete = index
			break
		}
	}

	// delete task
	tasks = append(tasks[:indexTaskToDelete], tasks[indexTaskToDelete+1:]...)

	// Save new task list
	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(FileName, jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The task has been deleted")
}

func ShowTask(FileName string) {
	var tasks []Task

	// Read from json
	data, err := os.ReadFile(FileName)
	if err != nil {
		log.Fatal(err)
	}

	// Deserializing JSON in a structure
	if err := json.Unmarshal(data, &tasks); err != nil {
		log.Fatal(err)
	}

	// Updating the last modified time
	now := time.Now().Format("02.01.2006 15:04:05")
	for i := range tasks {
		tasks[i].UpdateAt = now
	}

	DoneJSON, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(FileName, DoneJSON, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(DoneJSON))
}

func ShowDoneTask(FileName string) {
	var tasks []Task
	doneTask := []Task{}

	data, err := os.ReadFile(FileName)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now().Format("02.01.2006 15:04:05")
	for i := range tasks {
		tasks[i].UpdateAt = now
		if tasks[i].Status == "done" {
			doneTask = append(doneTask, tasks[i])
		}
	}

	DoneJSON, err := json.MarshalIndent(doneTask, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(DoneJSON))
}

func ShowTodoTask(FileName string) {
	var tasks []Task
	doneTask := []Task{}

	data, err := os.ReadFile(FileName)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now().Format("02.01.2006 15:04:05")
	for i := range tasks {
		tasks[i].UpdateAt = now
		if tasks[i].Status == "todo" {
			doneTask = append(doneTask, tasks[i])
		}
	}

	DoneJSON, err := json.MarshalIndent(doneTask, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(DoneJSON))
}

func ShowIdneDegressTask(FileName string) {
	var tasks []Task
	doneTask := []Task{}

	data, err := os.ReadFile(FileName)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now().Format("02.01.2006 15:04:05")
	for i := range tasks {
		tasks[i].UpdateAt = now
		if tasks[i].Status == "inde-degress" {
			doneTask = append(doneTask, tasks[i])
		}
	}

	DoneJSON, err := json.MarshalIndent(doneTask, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(DoneJSON))
}
