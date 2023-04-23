package main

import (
	"src/db"
	"src/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db.InitDB()
	defer db.CloseDB()

	router.GET("/books", handlers.GetBooks)
	router.GET("/books/:id", handlers.GetBook)
	router.POST("/books", handlers.AddBook)
	router.PUT("/books/:id", handlers.UpdateBook)
	router.DELETE("/books/:id", handlers.DeleteBook)

	router.Run(":8080")
}
