package courses

import (
	"backend-dev-path/initializers"
	"backend-dev-path/src/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)



func UpdateCourse(ctx *gin.Context) { 
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil { 
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id should be number",
		})
		return
	} 
	course := models.Courses{} 
	if err := initializers.DB.First(&course, id).Error; err != nil  {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
			"id": id,
		})
		return
	}

	
	var body struct { 
		Title       string `json:"title" binding:"required,min=3,max=100"`
		Description string `json:"description" binding:"omitempty,max=1000"`
		Category    string `json:"category" binding:"required"`
		Level       string `json:"level" binding:"required"`
		Price       float64    `json:"price" binding:"min=0"`
		Rating      float64    `json:"rating" binding:"min=0,max=5"`
		Duration    int    `json:"duration" binding:"min=0"`
		Instructor  string `json:"instructor" binding:"max=50"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil { 
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read JSON",
		})
		return
	}

	course.Title = body.Title
	course.Description = body.Description
	course.Category = body.Category
	course.Level = body.Level
	course.Price = body.Price
	course.Rating = body.Rating
	course.Duration = body.Duration
	course.Instructor = body.Instructor
	

	if err := initializers.DB.Save(&course).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update course",
		})
		return
	}
	ctx.JSON(http.StatusOK, course)

}