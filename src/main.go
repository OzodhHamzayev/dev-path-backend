package main

import (
	"backend-dev-path/initializers"
	"backend-dev-path/src/middleware"
	"backend-dev-path/src/modules/auth"
	"backend-dev-path/src/modules/courses"

	"github.com/gin-gonic/gin"
)


func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}


func main() { 
	
	r := gin.Default()
	r.Use(func(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
})
	admin := r.Group("/api/admin")
	admin.Use(middleware.CheckJWTToken)
	admin.Use(middleware.RequeireAdmin)

	r.POST("/api/auth/register", auth.Register)
	r.POST("/api/auth/login", auth.Login)

	admin.POST("/api/admin/courses", courses.CreateCourse)
	admin.DELETE("/api/admin/courses/:id", courses.DeleteCourse)
	admin.PUT("/api/admin/courses/:id", courses.UpdateCourse)
	admin.GET("/api/admin/courses", courses.GetAllCourses)
	r.Run()
}