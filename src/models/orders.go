package models

import "time"


type Orders struct {
	ID             uint
	UserID         uint
	Users 		   Users `gorm:"foreignKey:UserId"`
	CourseID       uint
	Courses		   Courses `gorm:"foreignKey:CourseId"`
	Amount         int64
	Currency       string
	Status         string
	ExpiresAt      time.Time
	PaidAt         *time.Time
	IdempotencyKey string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}