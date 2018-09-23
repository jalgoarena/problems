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

func init() {
	const (
		staticDir        = "./problems"
		problemsFileName = "problems.json"
	)

	box := packr.NewBox(staticDir)

	problemsJson, err := box.Open(problemsFileName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "opening problems.json file: %v\n", err.Error())
		os.Exit(1)
	}

	err = app.LoadProblems(problemsJson)

	if err != nil {
		fmt.Fprintf(os.Stderr, "loading problems.json file: %v\n", err.Error())
		os.Exit(1)
	}

	fmt.Println("Problems loaded successfully")
}

func main() {
	const defaultPort = "8080"

	router := SetupRouter()

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = defaultPort
	}

	router.Run(":" + port)
}
