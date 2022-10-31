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
}