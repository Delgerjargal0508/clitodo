package main

import ("fmt")

type Task struct {
	ID int
	Title string
	Completed bool
}

var tasks []Task
var nextID int = 1

func main(){
	for {
		fmt.Println("/nTodo App")
		fmt.Println("1. Ажил нэмэх")
		fmt.Println("2. Ажлаа солих")
		fmt.Println("3. Ажил устгах")
		fmt.Println("4. Ажил дууссан гэж хадгалах")
		fmt.Println("5. Ажлуудаа харах")
		fmt.Println("6. App-с гарах")
		fmt.Print("Сонго: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addTask()
		case 2:
			editTask()
		case 3:
			deleteTask()
		case 4:
			completeTask()
		case 5:
			listTask()
		case 6:
			fmt.Println("App-с гарч байна.")
			return
		default:
			fmt.Println("Буруу сонголт.")
		}
	}
}

func addTask(){
	fmt.Println("Хийх ажлаа оруулна уу:")
	var title string
	fmt.Scan(&title)
	task := Task{ID: nextID, Title: title, Completed: false}
	tasks = append(tasks, task)
	nextID++
	fmt.Println("Амжилттай нэмлээ")
}

func editTask(){
	fmt.Println("Засварлах ажлын ID-г оруул")
	var id int
	fmt.Scan(&id)
	for i, task := range tasks{
		if task.ID == id {
			fmt.Print("Шинэ ажлаа оруулна уу")
			var newTitle string
			fmt.Scan(&newTitle)
			tasks[i].Title = newTitle
			fmt.Println("Амжилттай заслаа.")
			return
		}
	}
	fmt.Println("Байхгүй байна")
}

func deleteTask(){
	fmt.Print("Устгах ажлын ID оруул")
	var id int
	fmt.Scan(&id)

	for i, task := range tasks{
		if task.ID == id{
			tasks = append(tasks[:i], tasks[i+1]...)
			fmt.Println("Устгалаа.")
			return
		}
	}
	fmt.Println("Байхгүй байна")
}

func completeTask(){
	fmt.Print("Дуугах ажлын ID оруул")
	var id int
	fmt.Scan(&id)

	for i, task := range tasks{
		if task.ID == id{
			tasks[i].Completed = true
			fmt.Println("Ажлыг дуусгалаа")
			return
		}
	}
	fmt.Println("Ажил байхгүй")
}

func listTask(){
	fmt.Println("\nTasks:")
	if len(tasks) == 0{
		fmt.Println("Ажил алга")
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
