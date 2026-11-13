package enrollments

import (
	"backend-dev-path/initializers"
	"backend-dev-path/src/models"
)





func CreateEnrollments(user_id, course_id uint) error {

	enrollment := models.Enrollments{
		UserID: user_id,
		CourseID: course_id,
	}

	result := initializers.DB.Create(&enrollment)

	if result.Error != nil {
		return result.Error
	}
	return nil
}