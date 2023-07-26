package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/leandro-hoenen/go-cicd-pipeline-test/handlers"
)

func SetRoutes(router *gin.Engine) {
	router.GET("/hello", handlers.HandleAdd)
}
