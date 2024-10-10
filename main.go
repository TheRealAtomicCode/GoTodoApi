package main

import (
	"fmt"
	"time"
)

type Todo struct {
	ID              int       `gorm:"primaryKey;column:Id"`
	TaskDescription string    `gorm:"not null;column:TaskDescription"`
	IsCompleted     bool      `gorm:"not null;default:false;column:IsCompleted"`
	CreatedAt       time.Time `gorm:"autoCreateTime;column:CreatedAt"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime;column:UpdatedAt"`
}

func main() {
	// Connection string
	db := initDb()

	var userInput string
	inUse := true

	fmt.Println("Would you like to add or view your todos?")

	for inUse {

		fmt.Printf("\nPress 1 to view your todos\n")
		fmt.Printf("Press 2 to add a todo\n")
		fmt.Printf("Press q to Exit \n \n")

		fmt.Scanln(&userInput)

		switch userInput {
		case "1":
			getTodos(db)
		case "2":
			addTodos(db)
		case "q":
			inUse = false
		default:
			fmt.Println("No such command")
		}
	}

}
