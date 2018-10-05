package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/jalgoarena/problems/pb"
	"github.com/jalgoarena/problems/problm"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
		gRPCAddr = flag.String("grpc", ":8081", "gRPC listen address")
	)

	flag.Parse()
	ctx := context.Background()
	srv := problm.NewService()
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
		log.Println("http:", *httpAddr)
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

		log.Println("grpc:", *gRPCAddr)
		handler := problm.NewGRPCServer(ctx, endpoints)
		gRPCServer := grpc.NewServer()
		pb.RegisterProblemsStoreServer(gRPCServer, handler)
		errChan <- gRPCServer.Serve(listener)
	}()

	log.Fatalln(<-errChan)
}
