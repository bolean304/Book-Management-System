package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"book-management-system/constants"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Fetch credentials from environment variables
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf(constants.MySqlDSNQuery, username, password, host, port, database)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	if err != nil {
		return errors.New(fmt.Sprintf("failed to connect to database: %w", err))
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return errors.New(fmt.Sprintf("failed to get underlying DB: %w", err))
	}

	// Connection Pool Setting
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	// Ensure the DB is available
	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("failed to ping the DB: %w", err)
	}

	log.Println("Successfully connected to database")
	return nil
}
