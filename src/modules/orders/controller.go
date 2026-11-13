package orders

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context) {
	var body struct {
		CourseID uint `json:"course_id" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "course_id should be positive number",
		})
		return
	}

	value, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	userID, ok := value.(uint)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	order, status, message := createOrderService(
		userID,
		body.CourseID,
		ctx.GetHeader("Idempotency-Key"),
	)

	if message != "" {
		ctx.JSON(status, gin.H{"error": message})
		return
	}

	ctx.JSON(status, gin.H{"order": order})
}