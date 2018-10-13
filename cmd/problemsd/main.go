package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/jalgoarena/problems/pb"
	"github.com/jalgoarena/problems/pkg/problm"

	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
		gRPCAddr = flag.String("grpc", ":8081", "gRPC listen address")
	)

	flag.Parse()

	ctx := context.Background()
	svc := setupService(logger)

	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	endpoints := problm.MakeServerEndpoints(svc, logger)

	// HTTP Transport
	go func() {
		logger.Log("http", *httpAddr)
		handler := problm.MakeHTTPHandler(endpoints, logger)
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

	logger.Log("terminated", <-errChan)
}

func setupService(logger log.Logger) *problm.ProblemsService {

	fieldKeys := []string{"method", "error"}

	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "problems_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "problems_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)

	svc := problm.NewService()
	svc = problm.LoggingMiddleware(logger)(svc)
	svc = problm.InstrumentingMiddleware(requestCount, requestLatency)(svc)

	return &svc
}
