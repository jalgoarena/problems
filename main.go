package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr"
	"github.com/jalgoarena/problems-store/app"
	"log"
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
		log.Fatalf("opening problems.json file: %v\n", err)
	}

	if err = app.LoadProblems(problemsJson); err != nil {
		log.Fatalf("loading problems.json file: %v\n", err)
	}

	problemsJson.Close()
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
