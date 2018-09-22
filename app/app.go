package app

import (
	"../domain"
	"github.com/gin-gonic/gin"
)

// curl -i http://localhost:8080/api/v1/problems
func GetProblems(c *gin.Context) {
	c.JSON(200, domain.Problems)
}

// curl -i http://localhost:8080/api/v1/problems/fib
func GetProblem(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, filter(domain.Problems, func(problem domain.Problem) bool {
		return problem.Id == id
	}))
}

func filter(problems []domain.Problem, f func(problem domain.Problem) bool) domain.Problem {
	for _, problem := range problems {
		if f(problem) {
			return problem
		}
	}

	var empty domain.Problem
	return empty
}
