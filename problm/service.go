package problm

import (
	"context"
	"errors"
	"github.com/jalgoarena/problems/pb"
)

type ProblemsService interface {
	FindById(ctx context.Context, problemId string) (*pb.Problem, error)
	FindAll(ctx context.Context) (*string, error)
}

func NewService() ProblemsService {
	return &problemsService{}
}

type problemsService struct{}

func (problemsService) FindById(_ context.Context, problemId string) (*pb.Problem, error) {

	problem, err := first(problems, func(problem *pb.Problem) bool {
		return problem.Id == problemId
	})

	return problem, err
}

var ErrEmpty = errors.New("not found")

func first(problems []*pb.Problem, f func(problem *pb.Problem) bool) (*pb.Problem, error) {
	for _, problem := range problems {
		if f(problem) {
			return problem, nil
		}
	}

	return nil, ErrEmpty
}

func (problemsService) FindAll(_ context.Context) (*string, error) {
	return rawProblems, nil
}

type problemRequest struct {
	ProblemId string `json:"problemId"`
}

type problemResponse struct {
	Problem *pb.Problem `json:"problem"`
	Err     string      `json:"err,omitempty"`
}

type problemsRequest struct{}

type problemsResponse struct {
	Problems *string `json:"problems"`
	Err      string  `json:"err,omitempty"`
}

var (
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)
