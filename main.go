package main

import (
	"fmt"

	"github.com/jc21/golang-sql-pq-test.git/pkg/database"
)

func main() {
	db := database.PostgresDB()
	// db should not be able to connect
	pingError := db.Ping()
	// should cause panic

	if pingError != nil {
		fmt.Printf("Database is not available: %+v\n", pingError)
	} else {
		fmt.Println("It's alive!")
	}
}