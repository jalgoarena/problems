package app

import (
	"github.com/gin-gonic/gin"
)

// curl -i http://localhost:8080/api/v1/problems
func GetProblems(c *gin.Context) {
	c.JSON(200, gin.H{"problemId": "fib"})
}
