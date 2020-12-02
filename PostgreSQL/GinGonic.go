package main

import (
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id     string `json:"id" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

var DB *gorm.DB
var books []Book

func FindBooks(c *gin.Context) {
	var books []Book
	DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func main() {
	// Dtabase connection
	db, err := gorm.Open("postgres", "user=postgres password=belikemee dbname=try port=5432 sslmode=disable")

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Book{})

	DB = db

	r := gin.Default()

	// Routes
	r.GET("/books", FindBooks)

	r.Run()
}
