package orders

import (
	"backend-dev-path/initializers"
	"backend-dev-path/src/models"
	"backend-dev-path/src/modules/enrollments"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ConfirmDemoOrder(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id should be number",
		})
		return
	}
	var order models.Orders
	err = initializers.DB.Where("id = ?", id).First(&order).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "order not found",
		})
		return
	}

	if err != nil { 
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Database error",
		})
		return
	}

	result := initializers.DB.
		Model(&models.Orders{}).
		Where("id = ?", id).
		Where("status = ?", "pending").
		Update("status", "paid")

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Database error",
		})
		return
	}
	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusConflict, gin.H{
			  "error": "Order is not pending or has already been confirmed",
		})
		return
	}
	exists, err := enrollments.CheckEnrollmentExists(order.UserID, order.CourseID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	if exists {
		ctx.JSON(http.StatusConflict, gin.H{
	       "error": "User is already enrolled in this course",
		})
		return
	}




	err = enrollments.CreateEnrollments(order.UserID, order.CourseID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":   "Enrollment created",
		"order_id":  order.ID,
		"user_id":   order.UserID,
		"course_id": order.CourseID,
		"status":    "paid",
	})

}