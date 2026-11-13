package orders

import (
	"backend-dev-path/src/models"
	"errors"
	"net/http"
	"time"

	"gorm.io/gorm"
)

func createOrderService(userID uint, courseID uint, key string) (models.Orders, int, string) {
	var empty models.Orders

	course, err := findCourseByID(courseID)
	if err != nil {
		return empty, http.StatusNotFound, "id not found"
	}

	if key == "" || len(key) > 100 {
		return empty, http.StatusBadRequest,
			"Idempotency-Key must contain 1–100 bytes"
	}

	var order models.Orders

	err = findOrderByKey(userID, key, &order)
	if err == nil {
		if order.CourseID != courseID {
			return empty, http.StatusConflict,
				"Idempotency-Key was used for another course"
		}

		return order, http.StatusOK, ""
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return empty, http.StatusInternalServerError,
			"Failed to check idempotency key"
	}

	err = findEnrollment(userID, courseID)
	if err == nil {
		return empty, http.StatusAlreadyReported,
			"You already have this course"
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return empty, http.StatusInternalServerError,
			"Failed to check enrollment"
	}

	err = expirePendingOrders(userID, courseID)
	if err != nil {
		return empty, http.StatusInternalServerError,
			"Failed to expire orders"
	}

	err = findPendingOrder(userID, courseID, &order)
	if err == nil {
		return order, http.StatusOK, ""
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return empty, http.StatusInternalServerError,
			"Failed to check existing order"
	}

	newOrder := models.Orders{
		UserID:         userID,
		CourseID:       courseID,
		Amount:         int64(course.Price),
		Currency:       "UZS",
		Status:         "pending",
		ExpiresAt:      time.Now().Add(30 * time.Minute),
		PaidAt:         nil,
		IdempotencyKey: key,
	}

	created := saveOrder(&newOrder)

	if created.Error != nil {
		return empty, http.StatusInternalServerError,
			"Failed to create order"
	}

	if created.RowsAffected == 0 {
		var existing models.Orders

		err := findOrderByKey(userID, key, &existing)
		if err != nil {
			return empty, http.StatusInternalServerError,
				"Failed to retrieve existing order"
		}

		if existing.CourseID != courseID {
			return empty, http.StatusConflict,
				"Idempotency-Key was used for another course"
		}

		return existing, http.StatusOK, ""
	}

	return newOrder, http.StatusCreated, ""
}