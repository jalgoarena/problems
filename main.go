package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.GET("/problems", GetProblems)
	}

	return router
}

func GetProblems(c *gin.Context) {
	fmt.Println("curl -i http://localhost:8080/api/v1/problems")
	c.JSON(200, gin.H{"problemId": "fib"})
}

func main() {
	router := SetupRouter()
	router.Run(":8080")
}
