package models



type Enrollments struct{ 
	ID       uint `gorm:"primaryKey" json:"id"`
	UserID   uint `json:"user_id"`
	CourseID uint `json:"course_id"`
}