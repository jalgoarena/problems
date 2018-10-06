package problm

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/jalgoarena/problems/pb"
	"time"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   ProblemsService
}

func (mw LoggingMiddleware) FindById(ctx context.Context, problemId string) (p *pb.Problem, err error) {
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

func (mw LoggingMiddleware) FindAll(ctx context.Context) (p *string, err error) {
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
