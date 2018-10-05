package problm

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/jalgoarena/problems/pb"
)

type Service interface {
	FindById(ctx context.Context, problemId string) (*pb.Problem, error)
	FindAll(ctx context.Context) (*string, error)
}

func NewService() Service {
	return problemsService{}
}

type problemsService struct{}

func (problemsService) FindById(_ context.Context, problemId string) (*pb.Problem, error) {

	problem := first(problems, func(problem *pb.Problem) bool {
		return problem.Id == problemId
	})

	return problem, nil
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

func MakeProblemEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(problemRequest)
		problem, err := svc.FindById(ctx, req.ProblemId)
		if err != nil {
			return problemResponse{nil, err.Error()}, nil
		}
		return problemResponse{Problem: problem, Err: ""}, nil
	}
}

func MakeProblemsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(problemsRequest)
		problems, err := svc.FindAll(ctx)
		if err != nil {
			return problemsResponse{nil, err.Error()}, nil
		}
		return problemsResponse{Problems: problems, Err: ""}, nil
	}
}

type Endpoints struct {
	ProblemEndpoint  endpoint.Endpoint
	ProblemsEndpoint endpoint.Endpoint
}

func (e Endpoints) FindById(ctx context.Context, problemId string) (*pb.Problem, error) {
	req := problemRequest{problemId}
	resp, err := e.ProblemEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}

	problemResp := resp.(problemResponse)
	if problemResp.Err != "" {
		return nil, errors.New(problemResp.Err)
	}
	return problemResp.Problem, nil
}

func (e Endpoints) FindAll(ctx context.Context) (*string, error) {
	req := problemsRequest{}
	resp, err := e.ProblemsEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}

	problemsResp := resp.(problemsResponse)
	if problemsResp.Err != "" {
		return nil, errors.New(problemsResp.Err)
	}
	return problemsResp.Problems, nil
}
