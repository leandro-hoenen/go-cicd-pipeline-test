package handlers

import (
	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/leandro-hoenen/go-cicd-pipeline-test/services"
)

type CompInput struct {
	ComponentOne int `json:"component_one" binding:"required"`
	ComponentTwo int `json:"component_two" binding:"required"`
}

type CompResult struct {
	InputComponents CompInput `json:"input_components"`
	Result          int       `json:"result"`
}

func HandleRandomAdd(c *gin.Context) {
	summandOne := rand.Intn(100)
	summandTwo := rand.Intn(100)
	sum := services.AddOperation(summandOne, summandTwo)

	input := CompInput{
		ComponentOne: summandOne,
		ComponentTwo: summandTwo,
	}

	r := CompResult{
		InputComponents: input,
		Result:          sum,
	}

	c.JSON(200, gin.H{
		"addResult": r,
	})
}

func HandleRandomSub(c *gin.Context) {
	minuend := rand.Intn(100)
	subtrahend := rand.Intn(100)
	difference := services.SubOperation(minuend, subtrahend)

	input := CompInput{
		ComponentOne: minuend,
		ComponentTwo: subtrahend,
	}

	r := CompResult{
		InputComponents: input,
		Result:          difference,
	}

	c.JSON(200, gin.H{
		"subResult": r,
	})
}

func HandleAdd(c *gin.Context) {
	var input CompInput

	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid JSON input",
		})
		return
	}

	sum := services.AddOperation(input.ComponentOne, input.ComponentTwo)

	r := CompResult{
		InputComponents: input,
		Result:          sum,
	}

	c.JSON(200, gin.H{
		"addResult": r,
	})

}

func HandleSub(c *gin.Context) {
	var input CompInput

	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid JSON input",
		})
		return
	}

	difference := services.SubOperation(input.ComponentOne, input.ComponentTwo)

	r := CompResult{
		InputComponents: input,
		Result:          difference,
	}

	c.JSON(200, gin.H{
		"subResult": r,
	})
}
