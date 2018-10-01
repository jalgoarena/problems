package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr"
	"github.com/jalgoarena/problems-store/app"
	"log"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", app.HealthCheck)
	v1 := router.Group("api/v1")
	{
		v1.GET("/problems", app.GetProblems)
		v1.GET("/problems/:id", app.GetProblem)
	}

	return router
}

const (
	staticDir        = "./problems"
	problemsFileName = "problems.json"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8080", "Port to listen on")
	flag.Parse()

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
	router := SetupRouter()
	router.Run(":" + port)
}
