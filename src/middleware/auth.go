package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)



func CheckJWTToken(ctx *gin.Context) { 
	authHeader := ctx.GetHeader("Authorization")


	if authHeader == "" { 
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Authorization token not found",
		})
		ctx.Abort()
		return
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" || parts[1] == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Authorization header must be: Bearer TOKEN",
		})
		ctx.Abort()
		return
	}

	secret := os.Getenv("SECRET")
	if len(secret) == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "JWT_SECRET is not configured",
		})
		ctx.Abort()
		return
	}
	tokenString := parts[1]

	token, err :=jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, 
			fmt.Errorf(
				"unexpected signing method: %s",
				token.Method.Alg(),
			)
		}
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Token is invalid or expired",
		})
		ctx.Abort()
		return
	}
	
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token claims",
		})
		ctx.Abort()
		return
	}


	userIDValue, exists := claims["user_id"]
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "User ID not found in token",
		})
		ctx.Abort()
		return
	}

	userID, ok := userIDValue.(float64)
	if !ok || userID <= 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid user ID in token",
		})
		ctx.Abort()
		return
	}

	ctx.Set("userID", uint(userID))

	ctx.Next()
}