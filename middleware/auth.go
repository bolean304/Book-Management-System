package middleware

import (
	"errors"
	"log"
	"net/http"

	"book-management-system/config"
	"book-management-system/models"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

/*
------------------------------------------------------------------------------------------
This function authenticates in incoming request.
-------------------------------------------------------------------------------------------
*/

func AuthenticationMiddleware(c *gin.Context) {
	// Get token from cookies
	tokenString, err := c.Cookie("UserAuthorizationCredentials")
	if err != nil {
		fmt.Println(tokenString)
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Please Login!"})
		c.Abort() // Stop further processing
		return
	}

	// Parse and verify the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the token signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		err = godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ERROR: " + err.Error()})
		c.Abort()
		return
	}

	// Extract the claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ERROR: Invalid claims"})
		c.Abort()
		return
	}
	fmt.Printf("claims : %v", claims)
	userID := claims["sub"].(float64)

	// Fetch the user from the database
	var user models.User
	log.Println("Querying DB")
	result := config.DB.Where("id=?", userID).First(&user)
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
	log.Printf("db user data : %v", user)

	// Attach the user to the context
	c.Set("user", user)

	// Continue with the next handler
	c.Next()
}
