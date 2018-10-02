package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jalgoarena/problems-store/problems"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFibProblem(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testRouter := SetupRouter()

	req, err := http.NewRequest("GET", "/api/v1/problems/fib", nil)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Errorf("GET /api/v1/problems/fib failed with error %d.", err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)

	if resp.Code != 200 {
		t.Errorf("GET /api/v1/problems/fib failed with error %d.", err)
	}

	jsonParser := json.NewDecoder(resp.Body)
	var fib problems.Problem

	if err = jsonParser.Decode(&fib); err != nil {
		t.Errorf("Could not parse fib problem")
	}

	if fib.Id != "fib" {
		t.Errorf("Fib problem was wrongly parsed")
	}
}

func TestGetAllProblems(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testRouter := SetupRouter()

	req, err := http.NewRequest("GET", "/api/v1/problems", nil)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Errorf("GET /api/v1/problems failed with error %d.", err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)

	if resp.Code != 200 {
		t.Errorf("GET /api/v1/problems failed with error %d.", err)
	}

	jsonParser := json.NewDecoder(resp.Body)
	var problemsList []problems.Problem

	if err = jsonParser.Decode(&problemsList); err != nil {
		t.Errorf("Could not parse problems list")
	}

	if len(problemsList) <= 0 {
		t.Errorf("No problems loaded")
	}
}
