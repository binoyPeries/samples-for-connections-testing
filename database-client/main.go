package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")      
	dbName := os.Getenv("DB_NAME")
	table := os.Getenv("DB_TABLE")     

	if table == "" {
		log.Fatal("Environment variable DB_TABLE is required")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)

	// Open database connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to open DB connection: %v", err)
	}
	defer db.Close()

	// Ping the database to verify connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	}

	log.Printf("Connected to MySQL database: %s", dbName)

	// Execute SELECT * FROM table
	query := fmt.Sprintf("SELECT * FROM %s", table)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}
	defer rows.Close()

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		log.Fatalf("Failed to get columns: %v", err)
	}
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Print each row
	for rows.Next() {
		err := rows.Scan(scanArgs...)
		if err != nil {
			log.Fatalf("Row scan failed: %v", err)
		}

		var rowStr string
		for i, col := range values {
			rowStr += fmt.Sprintf("%s: %s\t", columns[i], string(col))
		}
		log.Println(rowStr)
	}
	if err := rows.Err(); err != nil {
		log.Fatalf("Error reading rows: %v", err)
	}
}
