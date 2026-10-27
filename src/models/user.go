package models


type Users struct {
	ID 	 	 uint `gorm:"primaryKey" json:"id"`
	Name     string 
	Email    string `json:"email" gorm:"uniqueIndex;not null"`
	Password string
	Role  	 string
}