package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println(" Go and MySQL")

	db, err := sql.Open("mysql", "root:shivamansi@tcp(127.0.0.1:3306)/testdb")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO school VALUES('KK')")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	fmt.Println("Inserted!")

}
