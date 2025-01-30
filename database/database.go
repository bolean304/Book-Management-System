package database

import "gorm.io/gorm"

func initDB() {
	dsn := "host=localhost user=postgres password=yourpassword dbname=bookdb port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&User{}, &Book{}, &BorrowRecord{})
}
