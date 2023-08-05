package models

import "time"

type Post struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	Content   string
	UserID    uint
	User      User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
