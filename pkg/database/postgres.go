package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // blank import to get the postgres package
)

// PostgresDB Create postgres client
func PostgresDB() *sql.DB {
	var db *sql.DB

	psqlSource := fmt.Sprintf(
		"host='%s' port='%s' user='%s' password='%s' dbname='%s' sslmode='%s' fallback_application_name='%s'",
		"doesnt-exist.example.com",
		"5432",
		"postgres",
		"",
		"example-db",
		"false",
		"jc21-pq-example",
	)

	db, err := sql.Open("postgres", psqlSource)
	if err != nil {
		fmt.Printf("PostgresError: %+v", err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("PostgresError: %+v", err)
		return nil
	}

	return db
}
