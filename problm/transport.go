package problm

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/jalgoarena/problems/pb"
)

func MakeProblemEndpoint(svc *ProblemsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(problemRequest)
		problem, err := (*svc).FindById(ctx, req.ProblemId)
		if err != nil {
			return problemResponse{nil, err.Error()}, nil
		}
		return problemResponse{Problem: problem, Err: ""}, nil
	}
}

func MakeProblemsEndpoint(svc *ProblemsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(problemsRequest)
		problems, err := (*svc).FindAll(ctx)
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

func TransportLoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			logger.Log("msg", "calling endpoint")
			defer logger.Log("msg", "called endpoint")
			return next(ctx, request)
		}
	}
}
