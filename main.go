package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "time"

	_ "github.com/lib/pq"
)

const (
	user = "postgres"
	password = "1234"
	host = "localhost"
	port = 5432
	dbname = "todolist"
)

func main(){
	connstr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatalf("Failed to open connection: %v", err)
	}
	dbManager := NewDBManager(db)
	// _, err = dbManager.Create(&Todo{
	// 	title: "Learning Python",
	// 	description: "I should Python in a week if i can't, I wll run 2 kilometres",
	// 	assignee: "Alimadad Ismoilov",
	// 	status: true,
	// 	deadline: time.Date(2022, 11, 2, 13, 00, 00, 00, time.Local),
	// })
	// if err != nil {
	// 	log.Fatalf("Failed to create Todo: %v", err)
	// }
	// todo2, err := dbManager.Get(3)
	// if err != nil {
	// 	log.Fatalf("Failed to get Todo: %v", err)
	// }
	// fmt.Println(*todo2)
	// todo, err := dbManager.Update(&Todo{
	// 	id: 3,
	// 	title: "My Job",
	// 	description: "I hate my current job",
	// 	status: true,
	// 	assignee: "Ismoil Rustamov",
	// })
	// if err != nil {
	// 	log.Fatalf("Failed to Update Todo: %v", err)
	// }
	// fmt.Println(todo)
	// err = dbManager.Delete(4)
	// if err != nil {
	// 	log.Fatalf("Failed to delete Todo; %v", err)
	// } else {
	// 	fmt.Println("Succesfully deleted...")
	// }
	todos, err := dbManager.GetAll(&GetAllParam{
		limit: 20,
		page: 1,
		title: "Ghost",
	})
	if err != nil {
		log.Fatalf("Failed to Get All info: %v", err)
	}
	PrintTodos(todos)
}

func PrintTodos(todos []*Todo) {
	for _, todo := range todos {
		PrintTodo(todo)
	}
} 

func PrintTodo(todo *Todo) {
	fmt.Println("----------- Todos -------------")
	fmt.Printf("Id: %v\n", todo.id)
	fmt.Printf("Title: %v\n", todo.title)
	fmt.Printf("Description: %v\n", todo.description)
	fmt.Printf("Assignee: %v\n", todo.assignee)
	fmt.Printf("Status: %v\n", todo.status)
	fmt.Printf("Deadline %v\n", todo.deadline)
	fmt.Printf("Created At: %v\n", todo.created_at)
}