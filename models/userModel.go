package models

import (
	"time"
)

// equals
type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"unique"`
	Username  string `gorm:"unique"`
	Password  string
	Post      []Post
	CreatedAt time.Time
	UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index"`
}
