package courses

import (
	"backend-dev-path/initializers"
	"backend-dev-path/src/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)



func DeleteCourse(ctx *gin.Context) { 
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid course id",
		})
		return
	}
	result := initializers.DB.Delete(&models.Courses{}, id)
	if result.Error != nil { 
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}
	if result.RowsAffected == 0 {
	ctx.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
	return
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"message": "course deleted",
	})
}