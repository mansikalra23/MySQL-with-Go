package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   string
	Name string
}

var users []User

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

	stmt, err := db.Query("SELECT id, name from users")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	for stmt.Next() {
		var user User
		err := stmt.Scan(&user.Id, &user.Name)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}
	fmt.Println(users)
}
