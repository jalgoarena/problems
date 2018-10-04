package probls

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/jalgoarena/problems/pb"
)

type Service interface {
	FindById(ctx context.Context, problemId string) *pb.Problem
	FindAll(ctx context.Context) *string
}

type problemsService struct{}

func NewService() Service {
	return problemsService{}
}

func (problemsService) FindById(_ context.Context, problemId string) *pb.Problem {

	problem := first(problems, func(problem *pb.Problem) bool {
		return problem.Id == problemId
	})

	return problem
}

func (problemsService) FindAll(_ context.Context) *string {
	return rawProblems
}

type problemRequest struct {
	ProblemId string `json:"problemId"`
}

type problemResponse struct {
	Problem pb.Problem `json:"problem"`
}

type problemsRequest struct{}

type problemsResponse struct {
	Problems string `json:"problems"`
}

func makeProblemEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(problemRequest)
		problem := svc.FindById(ctx, req.ProblemId)
		return problemResponse{*problem}, nil
	}
}

func makeProblemsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(problemsRequest)
		problems := svc.FindAll(ctx)
		return problemsResponse{*problems}, nil
	}
}
