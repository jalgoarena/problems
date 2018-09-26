package app

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jalgoarena/problems-store/domain"
	"io"
)

var problems []domain.Problem

// curl -i http://localhost:8080/api/v1/problems
func GetProblems(c *gin.Context) {
	c.JSON(200, problems)
}

// curl -i http://localhost:8080/api/v1/problems/fib
func GetProblem(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, filter(problems, func(problem domain.Problem) bool {
		return problem.Id == id
	}))
}

func LoadProblems(problemsJson io.Reader) error {
	jsonParser := json.NewDecoder(problemsJson)

	if err := jsonParser.Decode(&problems); err != nil {
		return err
	}

	return nil
}

func filter(problems []domain.Problem, f func(problem domain.Problem) bool) domain.Problem {
	for _, problem := range problems {
		if f(problem) {
			return problem
		}
	}

	return domain.Problem{}
}
