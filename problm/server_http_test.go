package problm

import (
	"context"
	"encoding/json"
	"github.com/jalgoarena/problems/pb"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFibProblem(t *testing.T) {
	svc := NewService()
	endpoint := MakeProblemEndpoint(&svc)
	testHandler := MakeHTTPHandler(context.Background(), Endpoints{ProblemEndpoint: endpoint})

	req, _ := http.NewRequest("GET", "/api/v1/problems/fib", nil)
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	testHandler.ServeHTTP(resp, req)

	if resp.Code != 200 {
		t.Errorf("GET /api/v1/problems/fib failed with response code %d.", resp.Code)
	}

	jsonParser := json.NewDecoder(resp.Body)
	var fib *pb.Problem

	if err := jsonParser.Decode(&fib); err != nil {
		t.Errorf("Could not parse fib problem")
	}

	if fib.Id != "fib" {
		t.Errorf("Fib problem was wrongly parsed")
	}
}

func TestGetAllProblems(t *testing.T) {
	svc := NewService()
	endpoint := MakeProblemsEndpoint(&svc)
	testHandler := MakeHTTPHandler(context.Background(), Endpoints{ProblemsEndpoint: endpoint})

	req, _ := http.NewRequest("GET", "/api/v1/problems", nil)
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	testHandler.ServeHTTP(resp, req)

	if resp.Code != 200 {
		t.Errorf("GET /api/v1/problems failed with response code: %d", resp.Code)
	}

	jsonParser := json.NewDecoder(resp.Body)
	var problems []*pb.Problem

	if err := jsonParser.Decode(&problems); err != nil {
		t.Errorf("Could not parse problems list")
	}

	if len(problems) <= 0 {
		t.Errorf("No problems loaded")
	}
}

func BenchmarkGetFibProblem(b *testing.B) {
	svc := NewService()
	endpoint := MakeProblemEndpoint(&svc)
	testHandler := MakeHTTPHandler(context.Background(), Endpoints{ProblemEndpoint: endpoint})

	for i := 0; i <= b.N; i++ {
		req, _ := http.NewRequest("GET", "/api/v1/problems/fib", nil)
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		testHandler.ServeHTTP(resp, req)

		if resp.Code != 200 {
			b.Errorf("GET /api/v1/problems/fib failed with response code %d.", resp.Code)
		}
	}
}

func BenchmarkGetAllProblems(b *testing.B) {
	svc := NewService()
	endpoint := MakeProblemsEndpoint(&svc)
	testHandler := MakeHTTPHandler(context.Background(), Endpoints{ProblemsEndpoint: endpoint})

	for i := 0; i <= b.N; i++ {
		req, _ := http.NewRequest("GET", "/api/v1/problems", nil)
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		testHandler.ServeHTTP(resp, req)

		if resp.Code != 200 {
			b.Errorf("GET /api/v1/problems failed with response code: %d", resp.Code)
		}
	}
}
