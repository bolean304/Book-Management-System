package controllers

import (
	"book-management-system/config"
	"book-management-system/models"
	"book-management-system/utils"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

/*
------------------------------------------------------------------
This function handles the request for Registering the User.
-------------------------------------------------------------------
*/
func RegisterUser(c *gin.Context) {
	fmt.Println("Inside RegisterUser")
	var user models.User

	// Bind incoming JSON to user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Validate the user input (username uniqueness and other validations)
	if err := user.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the user's password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Fatal("Password hashing failed : %v", err)
	}
	log.Println("hashedPassword : %v", hashedPassword)
	user.Password = hashedPassword

	// Create the new user
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	// Return success response
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	log.Println("Inside login")
	var user models.User

	// Bind incoming JSON to user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}
	log.Printf("login data : %v", user)
	var dbData models.User
	log.Println("Querying DB")
	result := config.DB.Where("username=?", user.Username).First(&dbData)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Fatal("Record not found on query DB")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User"})
		return
	}
	if result.Error != nil {
		log.Fatal("Error on query DB")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Try Again"})
		return
	}
	log.Println("db user data : %v", dbData)
	err := bcrypt.CompareHashAndPassword([]byte(dbData.Password), []byte(user.Password))
	if err != nil {
		log.Println("err of bcrp :%v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User"})
		return
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Username,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // expires in 24 hours
	})
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	token, err := claims.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Try Again"})
		return
	}
	c.SetCookie("UserAuthorizationCredentials", token, int(time.Now().Add(time.Hour*24).Unix()), "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Login Successful"})

}
