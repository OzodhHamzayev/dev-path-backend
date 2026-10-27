package courses

import (
	"backend-dev-path/initializers"
	"backend-dev-path/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func CreateCourse(ctx *gin.Context) {
	var body struct { 
		Title       string 	   `json:"title" binding:"required,min=3,max=100"`
		Description string 	   `json:"description" binding:"omitempty,max=1000"`
		Category    string 	   `json:"category" binding:"required"`
		Level       string 	   `json:"level" binding:"required"`
		Price       float64    `json:"price" binding:"min=0"`
		Rating      float64    `json:"rating" binding:"min=0,max=5"`
		Duration    int        `json:"duration" binding:"min=0"`
		Instructor  string     `json:"instructor" binding:"max=50"`
	}


	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read JSON",
		})
		return
	}

	course := models.Courses{
		Title: body.Title,
		Description: body.Description,
		Category: body.Title,
		Level: body.Title,
		Price: body.Price,
		Rating: body.Rating,
		Duration: body.Duration,
		Instructor: body.Title,
	}

	if err := initializers.DB.Create(&course).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create",
		})
		return
	}
	
	ctx.JSON(http.StatusCreated, gin.H{
		"id": course.ID,
		"condition": "successfully created",
	})
}