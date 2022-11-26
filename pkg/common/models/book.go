package models

type Book struct {
	ID          int    `gorm:"primaryKey:auto_increment" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
}
