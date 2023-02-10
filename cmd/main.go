package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Start web application.")
	db, err := sql.Open("postgres", "postgresql://postgres:master@<database_ip>/todos?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
