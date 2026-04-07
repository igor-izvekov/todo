package models

import (
	"time"
)

type User struct {
	ID        int    `gorm:"primaryKey"` // Google ID (sub)
	Email     string `gorm:"uniqueIndex;not null"`
	Name      string
	AvatarURL string
	CreatedAt time.Time
	Tasks     []Task `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type Task struct {
	ID        int    `gorm:"primaryKey"`
	UserID    int    `gorm:"index;not null"`
	Title     string `gorm:"not null"`
	Completed bool   `gorm:"default:false"`
	CreatedAt time.Time
	User      User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}
