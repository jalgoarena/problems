package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/jalgoarena/problems/probls"
	"log"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", probls.HealthCheck)
	v1 := router.Group("api/v1")
	{
		v1.GET("/problems", probls.GetProblems)
		v1.GET("/problems/:id", probls.GetProblem)
	}

	return router
}

var port string

func init() {
	log.SetFlags(log.LstdFlags)
	flag.StringVar(&port, "port", "8080", "Port to listen on")
	flag.Parse()
}

func main() {
	router := setupRouter()
	router.Run(":" + port)
}
