package models

type BorrowRecord struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	UserID   uint   `gorm:"index" json:"user_id"`
	BookID   uint   `gorm:"index" json:"book_id"`
	BookName string `json:"book_name"`
}
