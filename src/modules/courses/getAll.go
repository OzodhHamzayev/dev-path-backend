package courses

import (
	"backend-dev-path/initializers"
	"backend-dev-path/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)



func GetAllCourses(ctx *gin.Context) { 
	courses := make([]models.Courses, 0)
	result := initializers.DB.Find(&courses)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Courses not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"courses": courses,
	})

}