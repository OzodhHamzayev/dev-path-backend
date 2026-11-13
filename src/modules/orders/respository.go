package orders

import (
	"backend-dev-path/initializers"
	"backend-dev-path/src/models"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func findCourseByID(courseID uint) (models.Courses, error) {
	var course models.Courses

	err := initializers.DB.
		Where("id = ?", courseID).
		First(&course).Error

	return course, err
}

func findOrderByKey(userID uint, key string, order *models.Orders) error {
	return initializers.DB.
		Where("user_id = ? AND idempotency_key = ?", userID, key).
		First(order).Error
}

func findEnrollment(userID, courseID uint) error {
	var enrollment models.Enrollments

	return initializers.DB.
		Where("user_id = ?", userID).
		Where("course_id = ?", courseID).
		First(&enrollment).Error
}

func expirePendingOrders(userID, courseID uint) error {
	return initializers.DB.
		Model(&models.Orders{}).
		Where("user_id = ?", userID).
		Where("course_id = ?", courseID).
		Where("status = ?", "pending").
		Where("expires_at <= ?", time.Now()).
		Update("status", "expired").Error
}

func findPendingOrder(userID uint, courseID uint, order *models.Orders) error {
	return initializers.DB.
		Where("user_id = ?", userID).
		Where("course_id = ?", courseID).
		Where("status = ?", "pending").
		Where("expires_at > ?", time.Now()).
		First(order).Error
}

func saveOrder(order *models.Orders) *gorm.DB {
	return initializers.DB.
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "user_id"},
				{Name: "idempotency_key"},
			},
			DoNothing: true,
		}).
		Create(order)
}