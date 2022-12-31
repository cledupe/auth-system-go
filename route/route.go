package route

import (
	"github.com/cledupe/jwt-auth/controllers"
	"github.com/cledupe/jwt-auth/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(routes *gin.Engine) {
	routes.POST("/signup", controllers.Signup)
	routes.POST("/login", controllers.Login)
	routes.GET("/me", middleware.TokenAuthMiddleware, controllers.UserInfo)
}
