package main

import (
	"time"

	math "github.com/leandro-hoenen/go-cicd-pipeline-test/math"
)

func main() {
	index := 1
	for {
		doubleInt := index * 2

		println(math.Add(index, doubleInt))
		println(math.Sub(doubleInt, index))

		time.Sleep(30 * time.Second)

		index++
	}
}
