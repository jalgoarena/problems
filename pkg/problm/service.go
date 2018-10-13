package problm

import (
	"context"
	"errors"
	"github.com/jalgoarena/problems/pb"
)

type ServiceMiddleware func(service ProblemsService) ProblemsService

type ProblemsService interface {
	FindById(ctx context.Context, problemId string) (*pb.Problem, error)
	FindAll(ctx context.Context) (*string, error)
	HealthCheck(ctx context.Context) (*pb.HealthCheckResponse, error)
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

func (problemsService) FindAll(_ context.Context) (*string, error) {
	return rawProblems, nil
}

func (problemsService) HealthCheck(_ context.Context) (*pb.HealthCheckResponse, error) {
	if problems == nil || len(problems) == 0 {
		return &pb.HealthCheckResponse{Up: false, ProblemCount: 0}, nil
	}

	return &pb.HealthCheckResponse{Up: true, ProblemCount: int32(len(problems))}, nil
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

type healthCheckRequest struct{}

type healthCheckResponse struct {
	HealthCheckResult *pb.HealthCheckResponse `json:"healthCheckResult"`
}

var (
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)
