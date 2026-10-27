package models


type Courses struct { 
	ID uint `gorm:"primaryKey" json:"id"`
	Title string
	Description string
	Category string
	Level string
	Price float64
	Rating float64
	Duration int
	Instructor string
}