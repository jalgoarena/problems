package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr"
	"github.com/jalgoarena/problems-store/app"
	"os"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.GET("/problems", app.GetProblems)
		v1.GET("/problems/:id", app.GetProblem)
	}

	return router
}

func main() {
	const (
		staticDir        = "./problems"
		problemsFileName = "problems.json"
		defaultPort      = "8080"
	)

	box := packr.NewBox(staticDir)

	problemsJson, err := box.Open(problemsFileName)

	if err != nil {
		fmt.Println("[err] opening problems.json file", err.Error())
		return
	}

	err = app.LoadProblems(problemsJson)

	if err != nil {
		fmt.Println("[err] loading problems.json file", err.Error())
		return
	}

	router := SetupRouter()

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = defaultPort
	}

	router.Run(":" + port)
}
