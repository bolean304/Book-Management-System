package models

import (
	"book-management-system/config"
	"fmt"
)

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique;not null" json:"username" validate:"required,min=3,max=20"`
	Password string `gorm:"not null:size=100" json:"password" validate:"required,min=6"`
}

// Custom validation function for unique username
func (user *User) Validate() error {
	// Validate struct tags
	if err := validate.Struct(user); err != nil {
		return err
	}

	// Check if the username is unique
	var existingUser User
	if err := config.DB.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		return fmt.Errorf("Username already exists")
	}

	return nil
}
