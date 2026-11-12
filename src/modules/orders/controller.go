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
	ctx.JSON(http.StatusOK, gin.H{
		"course_id": body.CourseID,
	})
}