package main

import (
	"backend-dev-path/initializers"
	"backend-dev-path/src/common/middleware"
	"backend-dev-path/src/modules/auth"
	"backend-dev-path/src/modules/courses"
	"backend-dev-path/src/modules/enrollments"
	"backend-dev-path/src/modules/orders"

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
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Idempotency-Key")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
	})


	r.POST("/api/auth/register", auth.Register)
	r.POST("/api/auth/login", auth.Login)
	r.GET("api//courses", courses.GetAllCourses)
	r.GET("/api/courses/:id", courses.GetCourseById) 	
	r.POST("/api/orders", middleware.CheckJWTToken, orders.CreateOrder)
	r.POST("/api/orders/:id/demo-confirm", middleware.CheckJWTToken, orders.ConfirmDemoOrder)
	r.GET("/api/me/enrollments", middleware.CheckJWTToken, enrollments.GetUserCourseByID)


	admin := r.Group("/api/admin")
	admin.Use(middleware.CheckJWTToken)
	admin.Use(middleware.RequireAdmin)

	admin.POST("/courses", courses.CreateCourse)
	admin.PUT("/courses/:id", courses.UpdateCourse)
	admin.GET("/courses", courses.GetAllCourses)
	admin.DELETE("/courses/:id", courses.DeleteCourse)
	r.Run()
}