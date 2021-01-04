// CRUD using Gorilla Mux, GORM and PostgreSQL

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Student struct {
	Rno   string `json:"rno" binding:"required"`
	Sname string `json:"sname" binding:"required"`
}

var DB *gorm.DB
var students []Student

func FindStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var students []Student
	DB.Find(&students)

	json.NewEncoder(w).Encode(&students)
}

func FindStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var student Student

	if err := DB.Where("rno = ?", params["rno"]).First(&student).Error; err != nil {
		fmt.Fprintf(w, "Record not found.")
		return
	}

	DB.First(&student, params["rno"])
	json.NewEncoder(w).Encode(&student)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var student Student
	json.NewDecoder(r.Body).Decode(&student)
	DB.Create(&student)
	json.NewEncoder(w).Encode(&student)

	fmt.Fprintf(w, "Record created.")
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var student Student
	if err := DB.Where("rno = ?", params["rno"]).First(&student).Error; err != nil {
		fmt.Fprintf(w, "Record not found.")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	Sname := keyVal["sname"]

	DB.Model(&Student{}).Where("rno = ?", params["rno"]).Update("sname", Sname)

	fmt.Fprintf(w, "Record updated.")

}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var student Student

	if err := DB.Where("rno = ?", params["rno"]).First(&student).Error; err != nil {
		fmt.Fprintf(w, "Record not found.")
		return
	}

	DB.Where("rno = ?", params["rno"]).Delete(&student)

	fmt.Fprintf(w, "Record deleted.")
}

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=try sslmode=disable password=belikemee")

	if err != nil {
		panic("Failed to connect database!")
	}
	fmt.Println("Successfully connected!")

	db.AutoMigrate(&Student{})

	DB = db

	r := mux.NewRouter()
	r.HandleFunc("/students", FindStudents).Methods("GET")
	r.HandleFunc("/students/{rno}", FindStudent).Methods("GET")
	r.HandleFunc("/students", CreateStudent).Methods("POST")
	r.HandleFunc("/students/{rno}", UpdateStudent).Methods("PUT")
	r.HandleFunc("/students/{rno}", DeleteStudent).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
