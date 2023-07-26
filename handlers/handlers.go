package handlers

import (
	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/leandro-hoenen/go-cicd-pipeline-test/services"
)

type addResult struct {
	SummandOne int
	SummandTwo int
	Sum        int
}

func HandleAdd(c *gin.Context) {
	summandOne := rand.Intn(100)
	summandTwo := rand.Intn(100)
	sum := services.AddOperation(summandOne, summandTwo)

	r := addResult{
		SummandOne: summandOne,
		SummandTwo: summandTwo,
		Sum:        sum,
	}

	c.JSON(200, gin.H{
		"addResult": r,
	})
}
