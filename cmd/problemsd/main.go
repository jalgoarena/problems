package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/jalgoarena/problems/pb"
	"github.com/jalgoarena/problems/problm"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
		gRPCAddr = flag.String("grpc", ":8081", "gRPC listen address")
	)

	flag.Parse()
	ctx := context.Background()
	var srv problm.ProblemsService
	srv = problm.NewService()
	srv = problm.LoggingMiddleware{Logger: logger, Next: srv}
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	problemEndpoint := problm.MakeProblemEndpoint(srv)
	problemsEndpoint := problm.MakeProblemsEndpoint(srv)
	endpoints := problm.Endpoints{
		ProblemEndpoint:  problemEndpoint,
		ProblemsEndpoint: problemsEndpoint,
	}

	// HTTP Transport
	go func() {
		logger.Log("http", *httpAddr)
		handler := problm.MakeHTTPHandler(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	// gRPC Transport
	go func() {
		listener, err := net.Listen("tcp", *gRPCAddr)
		if err != nil {
			errChan <- err
			return
		}

		logger.Log("grpc", *gRPCAddr)
		handler := problm.NewGRPCServer(ctx, endpoints)
		gRPCServer := grpc.NewServer()
		pb.RegisterProblemsStoreServer(gRPCServer, handler)
		errChan <- gRPCServer.Serve(listener)
	}()

	logger.Log(<-errChan)
}
