package controllers

import (
	"book-management-system/config"
	"book-management-system/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
------------------------------------------------------------------
This function borrows books to the user , which are added in database
-------------------------------------------------------------------
*/
func HnadleBorrowBooksRequest(c *gin.Context) {
	var request struct {
		UserID  uint   `json:"user_id"`
		BookIDs []uint `json:"book_ids"`
	}

	// Parse the request JSON
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Check if user exists
	var user models.User
	if err := config.DB.First(&user, request.UserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Check if all books exist
	var books []models.Book
	config.DB.Where("id IN (?)", request.BookIDs).Find(&books)

	if len(books) != len(request.BookIDs) {
		c.JSON(http.StatusNotFound, gin.H{"error": "One or more books not found"})
		return
	}

	// Create borrow records and collect borrowed book details
	var borrowedBooks []models.Book
	for _, book := range books {
		borrowRecord := models.BorrowRecord{UserID: request.UserID, BookID: book.ID, BookName: book.Title}
		config.DB.Create(&borrowRecord)
		borrowedBooks = append(borrowedBooks, book)
	}

	// Return borrowed book details
	c.JSON(http.StatusOK, gin.H{
		"message":        "Books borrowed successfully",
		"borrowed_books": borrowedBooks,
	})
}

/*
------------------------------------------------------------------
This function return the books ,which have borrowed user.
-------------------------------------------------------------------
*/
func HandleBorrwedBooksFetchRequest(c *gin.Context) {
	userid := c.Param("id")
	// Check if user exists
	var user models.User
	if err := config.DB.First(&user, userid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	// Check if all books exist
	var borrwoedbooks []models.BorrowRecord
	config.DB.Where("user_id=?", userid).Find(&borrwoedbooks)
	// Return borrowed book details
	c.JSON(http.StatusOK, gin.H{
		"borrowed_books": borrwoedbooks,
	})

}
