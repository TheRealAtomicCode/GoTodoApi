package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func initDb() *gorm.DB {
	connectionString := "sqlserver://sa:123123@10.100.235.11:1433?database=RandomStuff"

	db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("An error occurred while opening the database: ", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("An error occurred while getting the database: ", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("An error occurred while connecting to the database: ", err)
	}

	fmt.Println("You have successfully connected Go to MSSQL")

	return db
}
