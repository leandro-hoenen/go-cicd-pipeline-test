package handlers

import (
	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/leandro-hoenen/go-cicd-pipeline-test/services"
)

type compResult struct {
	ComponentOne int `json:"component_one"`
	ComponentTwo int `json:"component_two"`
	Result       int `json:"result"`
}

func HandleAdd(c *gin.Context) {
	summandOne := rand.Intn(100)
	summandTwo := rand.Intn(100)
	sum := services.AddOperation(summandOne, summandTwo)

	r := compResult{
		ComponentOne: summandOne,
		ComponentTwo: summandTwo,
		Result:       sum,
	}

	c.JSON(200, gin.H{
		"addResult": r,
	})
}

func HandleSub(c *gin.Context) {
	minuend := rand.Intn(100)
	subtrahend := rand.Intn(100)
	difference := services.SubOperation(minuend, subtrahend)

	r := compResult{
		ComponentOne: minuend,
		ComponentTwo: subtrahend,
		Result:       difference,
	}

	c.JSON(200, gin.H{
		"subResult": r,
	})
}
