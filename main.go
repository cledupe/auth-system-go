package main

import (
	"github.com/cledupe/jwt-auth/controllers"
	"github.com/cledupe/jwt-auth/initializers"
	"github.com/cledupe/jwt-auth/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDb()
	initializers.MigrateTables()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/me", middleware.TokenAuthMiddleware, controllers.UserInfo)
	r.Run()
}
