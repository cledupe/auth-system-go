package main

import (
	"github.com/cledupe/jwt-auth/infrastructure"
	"github.com/cledupe/jwt-auth/initializers"
	"github.com/cledupe/jwt-auth/route"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	infrastructure.ConnectDb()
	infrastructure.MigrateTables()
}

func main() {
	r := gin.Default()
	route.SetupRoutes(r)
	r.Run()
}
