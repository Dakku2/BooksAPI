package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"src/db"
	"src/models"
)

func GetBooks(c *gin.Context) {
	rows, err := db.GetDB().Query("SELECT * FROM books")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	books := []models.Book{}
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.PublishedYear); err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}

	c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book
	err := db.GetDB().QueryRow("SELECT * FROM books WHERE id=$1", id).Scan(&book.ID, &book.Title, &book.Author, &book.PublishedYear)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, book)
}

func AddBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var id int
	err := db.GetDB().QueryRow("INSERT INTO books (title, author, published_year) VALUES ($1, $2, $3) RETURNING id", book.Title, book.Author, book.PublishedYear).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}

	book.ID = id

	c.JSON(http.StatusCreated, book)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err := db.GetDB().Exec("UPDATE books SET title=$1, author=$2, published_year=$3 WHERE id=$4", book.Title, book.Author, book.PublishedYear, id)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{"message": fmt.Sprintf("Book %s successfully updated", id)})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	_, err := db.GetDB().Exec("DELETE FROM books WHERE id=$1", id)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{"message": fmt.Sprintf("Book %s successfully deleted", id)})
}
