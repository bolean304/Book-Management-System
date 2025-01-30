package routes

import (
	"book-management-system/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.Login)

	// Book management routes
	r.POST("/add-book", controllers.AddBook)
	r.GET("/fetch-books", controllers.ViewBooks)
	r.GET("/books/search", controllers.SearchBooks)
	r.PUT("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	// Borrowing routes
	r.POST("/book-borrow", controllers.HnadleBorrowBooksRequest)
	r.GET("/fetch-borrowed-books/:id", controllers.HandleBorrwedBooksFetchRequest)
}
