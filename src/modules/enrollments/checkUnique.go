package enrollments

import (
	"backend-dev-path/initializers"
	"backend-dev-path/src/models"
	"errors"

	"gorm.io/gorm"
)


func CheckEnrollmentExists(userID, courseID uint) (bool, error) {
    var enrollment models.Enrollments

    result := initializers.DB.
        Where("user_id = ?", userID).
        Where("course_id = ?", courseID).
        First(&enrollment)

    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
        return false, nil
    }
    if result.Error != nil {
        return false, result.Error
    }
    return true, nil
}