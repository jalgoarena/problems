package problm

import (
	"context"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/jalgoarena/problems/pb"
)

type grpcServer struct {
	problem  grpctransport.Handler
	problems grpctransport.Handler
}

func (s *grpcServer) FindById(ctx context.Context, r *pb.ProblemRequest) (*pb.ProblemResponse, error) {
	_, resp, err := s.problem.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ProblemResponse), nil
}

func (s *grpcServer) FindAll(ctx context.Context, r *pb.ProblemsRequest) (*pb.ProblemsResponse, error) {
	_, resp, err := s.problems.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ProblemsResponse), nil
}

func EncodeGRPCProblemRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(problemRequest)
	return &pb.ProblemRequest{ProblemId: req.ProblemId}, nil
}

func DecodeGRPCProblemRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.ProblemRequest)
	return problemRequest{req.ProblemId}, nil
}

func EncodeGRPCProblemResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(problemResponse)
	return &pb.ProblemResponse{Problem: res.Problem, Err: res.Err}, nil
}

func DecodeGRPCProblemResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.ProblemResponse)
	return problemResponse{Problem: res.Problem, Err: res.Err}, nil
}

func EncodeGRPCProblemsRequest(_ context.Context, r interface{}) (interface{}, error) {
	_ = r.(problemsRequest)
	return &pb.ProblemsRequest{}, nil
}

func DecodeGRPCProblemsRequest(_ context.Context, r interface{}) (interface{}, error) {
	_ = r.(*pb.ProblemsRequest)
	return problemsRequest{}, nil
}

func EncodeGRPCProblemsResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(problemsResponse)
	return &pb.ProblemsResponse{Problems: *res.Problems, Err: res.Err}, nil
}

func DecodeGRPCProblemsResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.ProblemsResponse)
	return problemsResponse{Problems: &res.Problems, Err: res.Err}, nil
}

func NewGRPCServer(ctx context.Context, endpoints Endpoints) pb.ProblemsStoreServer {
	return &grpcServer{
		problem: grpctransport.NewServer(
			endpoints.ProblemEndpoint,
			DecodeGRPCProblemRequest,
			EncodeGRPCProblemResponse,
		),
		problems: grpctransport.NewServer(
			endpoints.ProblemsEndpoint,
			DecodeGRPCProblemsRequest,
			EncodeGRPCProblemsResponse,
		),
	}
}
