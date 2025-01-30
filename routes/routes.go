package routes

import (
	"book-management-system/controllers"
	"book-management-system/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.Login)
	r.Use(middleware.AuthenticationMiddleware)
	// Book management routes
	r.POST("/add-book", controllers.HandleAddBook)
	r.GET("/fetch-books", controllers.HandleViewBooks)
	r.GET("/books/search", controllers.HandleSearchBooks)
	r.PUT("/books/:id", controllers.HandleUpdateBook)
	r.DELETE("/books/:id", controllers.HandleDeleteBook)

	// Borrowing routes
	r.POST("/book-borrow", controllers.HnadleBorrowBooksRequest)
	r.GET("/fetch-borrowed-books/:id", controllers.HandleBorrwedBooksFetchRequest)
}
