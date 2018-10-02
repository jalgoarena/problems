package problems

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

var problems []Problem

func HealthCheck(c *gin.Context) {
	if problems == nil || len(problems) == 0 {
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": "fail", "reason": "problems setup failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "problemsCount": len(problems)})
}

// curl -i http://localhost:8080/api/v1/problems
func GetProblems(c *gin.Context) {
	c.JSON(http.StatusOK, problems)
}

// curl -i http://localhost:8080/api/v1/problems/fib
func GetProblem(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, First(problems, func(problem Problem) bool {
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

func First(problems []Problem, f func(problem Problem) bool) Problem {
	for _, problem := range problems {
		if f(problem) {
			return problem
		}
	}

	return Problem{}
}
