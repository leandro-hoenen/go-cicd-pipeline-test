package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leandro-hoenen/go-cicd-pipeline-test/routes"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	routes.SetRoutes(router)

	router.Run(":8080")
}
