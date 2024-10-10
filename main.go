package main

import (
	"fmt"
	"time"
)

type Todo struct {
	Id              int       `gorm:primaryKey`
	TaskDescription string    `gorm:not null`
	IsComplete      bool      `gorm:not null;default:false`
	CreatedAt       time.Time `gorm:autoCreateTime`
	UpdatedAt       time.Time `gorm:autoUpdateTime`
}

func main() {
	// Connection string
	db := initDb()

	fmt.Println("Would you like to add or view your todos?")
	fmt.Println("Click 1 to view your todos")
	fmt.Println("Click 2 to add a todo")

	var userInput int = 0
	fmt.Scanln(userInput)

	fmt.Println(userInput)
	fmt.Println(db)
}
