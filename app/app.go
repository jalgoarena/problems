package app

import (
	"../domain"
	"github.com/gin-gonic/gin"
)

// curl -i http://localhost:8080/api/v1/problems
func GetProblems(c *gin.Context) {
	c.JSON(200, domain.Problems)
}
