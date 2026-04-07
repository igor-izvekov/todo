package models

import (
	"time"
)

type User struct {
	ID    int    `gorm:"primaryKey"`
	Email string `gorm:"uniqueIndex;not null"`
	Tasks []Task `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type Task struct {
	ID        int    `gorm:"primaryKey"`
	UserID    int    `gorm:"index;not null"`
	Title     string `gorm:"not null"`
	Completed bool   `gorm:"default:false"`
	CreatedAt time.Time
	User      User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}
