package enrollments

import (
	"backend-dev-path/initializers"
	"backend-dev-path/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)



func GetUserCourseByID(ctx *gin.Context) { 
	id, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "user or id not found",
		})
		ctx.Abort()
		return
	}

	var enrollments []models.Enrollments
	result := initializers.DB.Where("user_id = ?", id).Find(&enrollments)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error try again",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user's enrollments": enrollments,
	})
}