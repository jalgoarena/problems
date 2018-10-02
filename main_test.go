package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jalgoarena/problems-store/domain"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFibProblem(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testRouter := setupRouter()

	req, _ := http.NewRequest("GET", "/api/v1/problems/fib", nil)
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)

	if resp.Code != 200 {
		t.Errorf("GET /api/v1/problems/fib failed with response code %d.", resp.Code)
	}

	jsonParser := json.NewDecoder(resp.Body)
	var fib domain.Problem

	if err := jsonParser.Decode(&fib); err != nil {
		t.Errorf("Could not parse fib problem")
	}

	if fib.Id != "fib" {
		t.Errorf("Fib problem was wrongly parsed")
	}
}

func TestGetAllProblems(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testRouter := setupRouter()

	req, _ := http.NewRequest("GET", "/api/v1/problems", nil)
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)

	if resp.Code != 200 {
		t.Errorf("GET /api/v1/problems failed with response code: %d", resp.Code)
	}

	jsonParser := json.NewDecoder(resp.Body)
	var problems domain.Problems

	if err := jsonParser.Decode(&problems); err != nil {
		t.Errorf("Could not parse problems list")
	}

	if len(problems) <= 0 {
		t.Errorf("No problems loaded")
	}
}

func BenchmarkGetFibProblem(b *testing.B) {
	gin.SetMode(gin.TestMode)
	testRouter := setupRouter()

	for i := 0; i <= b.N; i++ {
		req, _ := http.NewRequest("GET", "/api/v1/problems/fib", nil)
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		testRouter.ServeHTTP(resp, req)

		if resp.Code != 200 {
			b.Errorf("GET /api/v1/problems/fib failed with response code %d.", resp.Code)
		}
	}
}

func BenchmarkGetAllProblems(b *testing.B) {
	gin.SetMode(gin.TestMode)
	testRouter := setupRouter()

	for i := 0; i <= b.N; i++ {
		req, _ := http.NewRequest("GET", "/api/v1/problems", nil)
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		testRouter.ServeHTTP(resp, req)

		if resp.Code != 200 {
			b.Errorf("GET /api/v1/problems failed with response code: %d", resp.Code)
		}
	}
}
