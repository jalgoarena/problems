package grpc

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/jalgoarena/problems/pb"
	"github.com/jalgoarena/problems/problm"
	"google.golang.org/grpc"
)

func New(conn *grpc.ClientConn) problm.ProblemsService {
	problemEndpoint := grpctransport.NewClient(
		conn, "pb.ProblemsStore", "FindById",
		problm.EncodeGRPCProblemRequest,
		problm.DecodeGRPCProblemResponse,
		pb.ProblemResponse{},
	).Endpoint()
	problemsEndpoint := grpctransport.NewClient(
		conn, "pb.ProblemsStore", "FindAll",
		problm.EncodeGRPCProblemsRequest,
		problm.DecodeGRPCProblemsResponse,
		pb.ProblemsResponse{},
	).Endpoint()
	return problm.Endpoints{
		ProblemEndpoint:  problemEndpoint,
		ProblemsEndpoint: problemsEndpoint,
	}
}
