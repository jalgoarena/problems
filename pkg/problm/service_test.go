package problm

import (
	"context"
	"encoding/json"
	"testing"
)

func TestGetProblem(t *testing.T) {
	srv := NewService()
	ctx := context.Background()

	problem, err := srv.FindById(ctx, "fib")
	if err != nil {
		t.Error(err)
	}

	if problem == nil {
		t.Errorf("FindById(%q) is nil", "fib")
	}

	if problem.Id != "fib" {
		t.Errorf("FindById(%q).Id != %q", "fib", problem.Id)
	}
}

func TestGetProblems(t *testing.T) {
	srv := NewService()
	ctx := context.Background()

	problemsJSON, err := srv.FindAll(ctx)
	if err != nil {
		t.Error(err)
	}

	if problemsJSON == nil {
		t.Error("FindAll() is nil")
	}

	bytes := []byte(*problemsJSON)
	var problems Problems

	if err := json.Unmarshal(bytes, &problems); err != nil {
		t.Errorf("json.Unmarshall of FindAll() result: %v", err)
	}

	if len(problems) < 100 {
		t.Errorf("len(FindAll()) < 100, count: %d", len(problems))
	}

	problem := problems.First(func(problem *Problem) bool {
		return problem.Id == "fib"
	})

	if problem.Id != "fib" {
		t.Error("FindAll()['fib'].Id != 'fib'")
	}
}
