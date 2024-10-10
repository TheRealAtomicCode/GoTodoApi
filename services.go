package main

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

func addTodos(db *gorm.DB) {

	var description string = ""
	fmt.Println("Please enter the task description:")

	fmt.Scanln(&description)

	newTodo := Todo{TaskDescription: description}

	result := db.Create(&newTodo)

	if result.Error != nil {
		log.Fatal("Failed to add todo: ", result.Error)
	}

	fmt.Println("Added the todo")
	fmt.Println(newTodo.ID)

}

func getTodos(db *gorm.DB) {

	var todos []Todo

	result := db.Find(&todos)

	if result.Error != nil {
		log.Println("Failed to get totos: ", result.Error)
	}

	fmt.Println("The todos:")
	for _, todo := range todos {
		fmt.Printf("ID: %d, description: %s, completed: %v \n", todo.ID, todo.TaskDescription, todo.IsCompleted)
	}

}
