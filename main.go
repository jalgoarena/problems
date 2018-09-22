package main

import (
	"./app"
	"./domain"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr"
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
	box := packr.NewBox("./problems")

	problemsJson, err := box.Open("problems.json")
	if err != nil {
		fmt.Println("[err] opening problems.json file: ", err.Error())
		return
	}

	jsonParser := json.NewDecoder(problemsJson)

	if err = jsonParser.Decode(&domain.Problems); err != nil {
		fmt.Println("[err] parsing problems.json file", err.Error())
		return
	}

	router := SetupRouter()
	router.Run(":8080")
}
