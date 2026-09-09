package models

import (
	"time"
)

// Post merepresentasikan tabel 'posts' di database
type Post struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"type:varchar(200)" json:"title"`
	Content     string    `gorm:"type:text" json:"content"`
	Category    string    `gorm:"type:varchar(100)" json:"category"`
	CreatedDate time.Time `gorm:"autoCreateTime;column:created_date" json:"created_date"`
	UpdatedDate time.Time `gorm:"autoUpdateTime;column:updated_date" json:"updated_date"`
	Status      string    `gorm:"type:varchar(100)" json:"status"`
}

// ArticleRequest merepresentasikan payload untuk Create dan Update
type ArticleRequest struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Status   string `json:"status"`
}
