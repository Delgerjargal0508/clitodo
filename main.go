package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
type Task struct {
	ID        int
	Title     string
	Completed bool
}

var tasks []Task
var nextID int = 1

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nTo-do App")
		fmt.Println("1. Task нэмэх")
		fmt.Println("2. Task шинэчлэх")
		fmt.Println("3. Task устгах")
		fmt.Println("4. Task дуусгах")
		fmt.Println("5. Task харах")
		fmt.Println("6. App-с гарах")
		fmt.Print("Сонголт: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addTask(reader)
			listTask()
		case 2:
			editTask(reader)
			listTask()
		case 3:
			deleteTask()
			listTask()
		case 4:
			completeTask()
			listTask()
		case 5:
			listTask()
		case 6:
			fmt.Println("App-с гарч байна.")
			return
		default:
			fmt.Println("Буруу сонголт. Дахин оролдоно уу.")
		}
	}
}

func addTask(reader *bufio.Reader) {
	fmt.Println("Нэмэх Task-аа бичээрэй:")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)
	task := Task{ID: nextID, Title: title, Completed: false}
	tasks = append(tasks, task)
	nextID++
	fmt.Println("Амжилттай нэмлээ.")
}

func editTask(reader *bufio.Reader) {
	fmt.Println("Шинэчлэх task-н ID-г оруул:")
	var id int
	fmt.Scanln(&id)

	for i, task := range tasks {
		if task.ID == id {
			fmt.Print("Шинэ task оруулаарай: ")
			newTitle, _ := reader.ReadString('\n')
			newTitle = strings.TrimSpace(newTitle)
			tasks[i].Title = newTitle
			fmt.Println("Амжилттай шинэчиллээ.")
			return
		}
	}
	fmt.Println("ID олдсонгүй.")
}

func deleteTask() {
	fmt.Print("Устгах task-н ID оруулаарай: ")
	var id int
	fmt.Scanln(&id)

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Устгагдлаа.")
			return
		}
	}
	fmt.Println("ID олдсонгүй.")
}

func completeTask() {
	fmt.Print("Дуусгах task-н ID оруул: ")
	var id int
	fmt.Scanln(&id)

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
			fmt.Println("Task дуусгалаа.")
			return
		}
	}
	fmt.Println("ID олдсонгүй.")
}

func listTask() {
	fmt.Println("\nTasks:")
	if len(tasks) == 0 {
		fmt.Println("Ажил алга.")
		return
	}
	for _, task := range tasks {
		status := "Дуусаагүй"
		if task.Completed {
			status = "Дууссан"
		}
		fmt.Printf("ID: %d | Title: %s | Status: %s\n", task.ID, task.Title, status)
	}
}
