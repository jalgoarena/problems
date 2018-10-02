package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr"
	"log"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", HealthCheck)
	v1 := router.Group("api/v1")
	{
		v1.GET("/problems", GetProblems)
		v1.GET("/problems/:id", GetProblem)
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
	log.SetFlags(log.LstdFlags)

	box := packr.NewBox(staticDir)
	problemsJson, err := box.Open(problemsFileName)

	if err != nil {
		log.Fatalf("opening problems.json file: %v\n", err)
	}

	if err = LoadProblems(problemsJson); err != nil {
		log.Fatalf("loading problems.json file: %v\n", err)
	}

	problemsJson.Close()
	log.Println("Problems loaded successfully")
}

func main() {
	router := SetupRouter()
	router.Run(":" + port)
}
