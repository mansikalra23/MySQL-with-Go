package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:belikemee@tcp(127.0.0.1:3306)/try")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	stmt, err := db.Prepare("INSERT INTO users (id, name) VALUES (?, ?)")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec("2", "b")
	if err != nil {
		panic(err)
	}
	fmt.Println("New record entered.")
}
