package main

import (
	"context"
	"flag"
	"fmt"
	grpcclient "github.com/jalgoarena/problems/client/grpc"
	"github.com/jalgoarena/problems/pkg/problm"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	var (
		gRPCAddr = flag.String("addr", ":8081", "gRPC address")
	)
	flag.Parse()
	ctx := context.Background()
	conn, err := grpc.Dial(*gRPCAddr, grpc.WithInsecure(), grpc.WithTimeout(1*time.Second))
	if err != nil {
		log.Fatalln("gRPC dial:", err)
	}
	defer conn.Close()
	problemsService := grpcclient.New(conn)
	args := flag.Args()
	var cmd string
	cmd, args = pop(args)
	switch cmd {
	case "problem":
		var problemId string
		problemId, args = pop(args)
		problem(ctx, problemsService, problemId)
	case "problems":
		problems(ctx, problemsService)
	default:
		log.Fatalln("unknown command", cmd)
	}
}

func problem(ctx context.Context, service problm.ProblemsService, problemId string) {
	p, err := service.FindById(ctx, problemId)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(p)
}

func problems(ctx context.Context, service problm.ProblemsService) {
	p, err := service.FindAll(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(*p)
}

func pop(s []string) (string, []string) {
	if len(s) == 0 {
		return "", s
	}
	return s[0], s[1:]
}
