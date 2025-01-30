package controllers

import (
	"book-management-system/config"
	"book-management-system/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

/*
------------------------------------------------------------------
This function handles the request for adding the book from user.
-------------------------------------------------------------------
*/
func AddBook(c *gin.Context) {
	var book models.Book

	// Bind incoming JSON to book struct
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Validate the book (including checking for duplicates)
	if err := book.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the new book in the database
	if err := config.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add book"})
		return
	}

	// Return success response with the created book
	c.JSON(http.StatusCreated, gin.H{"message": "Book added successfully", "book": book})
}

/*
------------------------------------------------------------------------------------------
This function handles the request for fetching the all saved books in database to show on UI.
-------------------------------------------------------------------------------------------
*/

func ViewBooks(c *gin.Context) {
	var books []models.Book
	config.DB.Find(&books)
	c.JSON(http.StatusOK, books)
}

/*
------------------------------------------------------------------------------------------
This function handles the request for searching the book based on title and author.
-------------------------------------------------------------------------------------------
*/

func SearchBooks(c *gin.Context) {
	// Get the title and author query parameters from the request
	title := c.DefaultQuery("title", "")
	author := c.DefaultQuery("author", "")

	// Declare a slice to store the search results
	var books []models.Book

	// If neither title nor author is provided, return an error
	if title == "" && author == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "At least one of 'title' or 'author' query parameters is required"})
		return
	}

	// Build the query based on the provided parameters
	query := config.DB.Model(&models.Book{})
	if title != "" {
		query = query.Where("LOWER(title) LIKE ?", "%"+strings.ToLower(title)+"%")
	}
	if author != "" {
		query = query.Where("LOWER(author) LIKE ?", "%"+strings.ToLower(author)+"%")
	}

	// Execute the query
	err := query.Find(&books).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error searching books"})
		return
	}

	// Return the found books as a JSON response
	if len(books) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No books found"})
		return
	}

	c.JSON(http.StatusOK, books)
}

/*
------------------------------------------------------------------------------------------
This function handles the request for updating the book details based on ID.
-------------------------------------------------------------------------------------------
*/
func UpdateBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")

	if err := config.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	if err := book.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&book)
	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
}

/*
------------------------------------------------------------------------------------------
This function handles the request for deleting  the book details based on ID.
-------------------------------------------------------------------------------------------
*/
func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	config.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
