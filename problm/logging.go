package problm

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/jalgoarena/problems/pb"
	"time"
)

type loggingMiddleware struct {
	Next   ProblemsService
	Logger log.Logger
}

func LoggingMiddleware(logger log.Logger) ServiceMiddleware {
	return func(next ProblemsService) ProblemsService {
		return loggingMiddleware{next, logger}
	}
}

func (mw loggingMiddleware) FindById(ctx context.Context, problemId string) (p *pb.Problem, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", fmt.Sprintf("FindById('%s')", problemId),
			"output.Title", p.Title,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	p, err = mw.Next.FindById(ctx, problemId)
	return
}

func (mw loggingMiddleware) FindAll(ctx context.Context) (p *string, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "FindAll",
			"len(output)", len(*p),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	p, err = mw.Next.FindAll(ctx)
	return
}

func (mw loggingMiddleware) HealthCheck(ctx context.Context) (output *pb.HealthCheckResponse, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "HealthCheck",
			"up", output.Up,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.HealthCheck(ctx)
	return
}
