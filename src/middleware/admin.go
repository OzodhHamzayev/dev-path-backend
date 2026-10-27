package middleware

import (
	"backend-dev-path/initializers"
	"backend-dev-path/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func RequireAdmin(ctx *gin.Context) { 
	id, exists := ctx.Get("userID")
	if !exists {
	    ctx.JSON(http.StatusUnauthorized, gin.H{
            "error": "User ID not found",
        })
		ctx.Abort()
		return
	}

	user := models.Users{}
	result := initializers.DB.First(&user, id)
	if result.Error != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "User ID not found",
		})
		ctx.Abort()
		return
	}

	if user.Role != "admin" {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": "Admin access required",
		})
		ctx.Abort()
		return
	}

	ctx.Next()

}