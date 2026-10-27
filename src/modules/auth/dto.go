package auth

import (
	"backend-dev-path/initializers"
	"backend-dev-path/src/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"

)


func Register(ctx *gin.Context) { 
	var body struct { 
		Name     string `json:"name" binding:"required,min=2,max=50"`
		Email    string `json:"email" binding:"required,email,max=100"`
		Password string `json:"password" binding:"required,min=6,max=72"`
		Role 	 string `json:"role"`
	}

	
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read json",
		})
		return
	}
	user := models.Users{}
	err := initializers.DB.First(&user, "email = ?", body.Email).Error
	if err == nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"error": "The user already taken",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 12)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash the password",
		})
		return
	}

	user = models.Users{
		Name: body.Name,
		Email: body.Email,
		Password: string(hash),
		Role: "user",
	}

	if err := initializers.DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"user": "successfully created",
	})
}

func Login(ctx *gin.Context) { 
	var body struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read JSON",
		})
		return
	}

	user := models.Users{}
	result := initializers.DB.First(&user, "email = ?", body.Email).Error 
	if result != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "You have to register first",
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}


	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"name":  user.Name,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		"email": user.Email,
	})
	
	token, err := jwtToken.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"id": user.ID,
			"name": user.Name,
			"role": user.Role,
		},
	})

}