package main

import ("fmt")

type Task struct {
	ID        int
	Title     string
	Completed bool
}
//Taskuudiig dynamic arrayd hadgalj ID ugnu.
var tasks []Task
var nextID int = 1

//main functioniig for loop ashiglaj urgelj prompt hiine. fmtScan ashiglaj songolttoi hargalzah functioniig duudna
func main() {
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
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addTask()
			listTask()
		case 2:
			editTask()
			listTask()
		case 3:
			deleteTask()
			listTask()
		case 4:
			completeTask()
			listTask()
		case 5:
			fmt.Println("App-с гарч байна.")
			return
		default:
			fmt.Println("Буруу сонголт.")
		}
	}
}

func addTask() {
	fmt.Println("Нэмэх Task-аа бичээрэй:")
	var title string
	fmt.Scan(&title)
	task := Task{ID: nextID, Title: title, Completed: false}
	tasks = append(tasks, task)
	nextID++
	fmt.Println("Амжилттай нэмлээ")
}

func editTask() {
	fmt.Println("Шинэчлэх task-н ID-г оруул:")
	var id int
	fmt.Scan(&id)
	for i, task := range tasks {
		if task.ID == id {
			fmt.Print("Шинэ task оруулаарай: ")
			var newTitle string
			fmt.Scan(&newTitle)
			tasks[i].Title = newTitle
			fmt.Println("Амжилттай шинэчиллээ.")
			return
		}
	}
}

func deleteTask() {
	fmt.Print("Устгах task-н ID оруулаарай: ")
	var id int
	fmt.Scan(&id)

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Устгагдлаа.")
			return
		}
	}
}

func completeTask() {
	fmt.Print("Дуусгах task-н ID оруул: ")
	var id int
	fmt.Scan(&id)

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
			return
		}
	}
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
