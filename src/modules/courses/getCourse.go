package courses

import (
	"backend-dev-path/initializers"
	"backend-dev-path/src/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func GetCourseById(ctx *gin.Context) { 
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id should be number",
		})
		return
	}

	var course models.Courses
	err = initializers.DB.Where("id = ?", id).First(&course).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Course not found",
		})
		return
	} 
	
	ctx.JSON(http.StatusOK, gin.H{
		"course": course,
	})
}