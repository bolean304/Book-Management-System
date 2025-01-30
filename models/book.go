package models

import (
	"book-management-system/config"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Book struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	Title         string `gorm:"not null;index:idx_title_author_genre" json:"title" validate:"required,min=3,max=255"`
	Author        string `gorm:"not null;index:idx_title_author_genre" json:"author" validate:"required,min=3,max=255"`
	Genre         string `json:"genre" validate:"required,min=3,max=50"`
	PublishedYear int    `json:"published_year" validate:"required"`
}

var validate = validator.New()

// Validate method checks if the book is valid, including checking for duplicate books
func (book *Book) Validate() error {
	// Validate the fields based on tags
	if err := validate.Struct(book); err != nil {
		return err
	}

	// Check if the book already exists in the database based on Title, Author, and Genre
	var existingBook Book
	if err := config.DB.Where("title = ? AND author = ? AND genre = ?", book.Title, book.Author, book.Genre).First(&existingBook).Error; err == nil {
		return fmt.Errorf("duplicate book found: a book with the same title, author, and genre already exists")
	}

	return nil
}
