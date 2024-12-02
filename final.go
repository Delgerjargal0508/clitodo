package main

import ("fmt")
//Task struct uusgej ID ni taskiin daraalal, title ni task string, completed
type Task struct {
	ID 			int
	Title 		string
	Completed 	bool
}
// tasks dynamic array zarlasan
var tasks []Task
// arrayiin id-g newId variable zarlaj uguh
var newId int = 1
//main function. for loop hiij oroh uyes ehleed buh songoltiig consoled printlene. scan ashiglaj ali 1 case iig songohod holbogdoh function duudna
func main(){
	for {
		fmt.Println("/nTo-do App")
		fmt.Println("1. Task нэмэх")
		fmt.Println("2. Task өөрчлөх")
		fmt.Println("3. Task устгах")
		fmt.Println("4. Task дуусгах")
		fmt.Println("5. Task харах")
		fmt.Println("Аппаас гарах бол компьютерийн гараас Ctrl товч болон C товчийг зэрэг дар")

		var choice int
		fmt.Scan(&choice)
		switch choice{
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
		default:
			return
			fmt.Println("1-5 хооронд сонго")
		}
	}
}
//case 1 function. task arrayd nemne
func addTask(){
	fmt.Println("Нэмэх task-аа бич:")
	var title string
	fmt.Scan(&title)
	task := Task(ID: nextID, Title: title, Completed: false)
	tasks = append(tasks, task)
	nextID++
	fmt.Println("Task нэмэгдлээ.")
}
//case 2 edit hih function. ID oruulah promptoor damjuulad id-rn edit hiine
func editTask(){
	fmt.Println("Edit хийх Task-ийнхаа ID-г сонго:")
	var id int
	fmt.Scan(&id)
	for i, task := range tasks{
		if task.ID == id{
			fmt.Print("Шинээр хийх ажлаа оруул:")
			var newTitle string
			fmt.Scan(&newTitle)
			tasks.[i].Title = newTitle
			fmt.Println("Edit хийлээ.")
			return
		}
	}
}

func deleteTask(){
	fmt.Println("Устгах task-ийнхаа ID-г сонго:")
	var id int
	fmt.Scan(&id)
	for i, task := range tasks{
		if task.ID == id{
			tasks.append(tasks[:i], tasks[i+1;]...)
			fmt.Println("Устгасан")
			return
		}
	}
}

func completeTask(){
	fmt.Println("Дуусгасан ажлын ID Оруул:")
	var id int
	fmt.Scan(&id)
	for i, task := range tasks{
		if task.ID == id{
			tasks[i].Completed = true
			fmt.Printnl("Ажлыг дуусгаснаар бүртгэлээ.")
			return
		}
	}
}

func listTask(){
	fmt.Println("\nTasks:")
	if len(tasks) == 0{
		fmt.Println("Ажил алга.")
		return
	}
	for _, task := range tasks{
		status := "ДУусаагүй"
		if task.completed{
			status = "Дууссан"
		}
		fmt.Printf("ID: %d | Title %s | Status: %s\n", task.ID, task.Title, status)
	}
}