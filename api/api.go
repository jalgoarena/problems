package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr"
	"github.com/jalgoarena/problems-store/domain"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var problems domain.Problems
var rawProblems string

func init() {
	log.SetFlags(log.LstdFlags)
	box := packr.NewBox(".")
	problemsJSON, err := box.Open("problems.json")
	defer problemsJSON.Close()

	if err != nil {
		log.Fatalf("opening problems.json file: %v\n", err)
	}

	if err = loadProblems(problemsJSON); err != nil {
		log.Fatalf("loading problems.json file: %v\n", err)
	}

	log.Println("Problems loaded successfully")
}

func loadProblems(problemsJSON io.Reader) error {
	bytes, err := ioutil.ReadAll(problemsJSON)
	if err != nil {
		return err
	}

	rawProblems = string(bytes[:])

	if err := json.Unmarshal(bytes, &problems); err != nil {
		return err
	}

	return nil
}

// curl -i http://localhost:8080/health
func HealthCheck(c *gin.Context) {
	if problems == nil || len(problems) == 0 {
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": "fail", "reason": "problems setup failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "problemsCount": len(problems)})
}

// curl -i http://localhost:8080/api/v1/problems
func GetProblems(c *gin.Context) {
	c.String(http.StatusOK, rawProblems)
}

// curl -i http://localhost:8080/api/v1/problems/fib
func GetProblem(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, problems.First(func(problem domain.Problem) bool {
		return problem.Id == id
	}))
}
