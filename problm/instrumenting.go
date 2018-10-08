package problm

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/metrics"
	"github.com/jalgoarena/problems/pb"
	"time"
)

type InstrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	Next           ProblemsService
}

func (mw InstrumentingMiddleware) FindById(ctx context.Context, problemId string) (output *pb.Problem, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "findbyid", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.FindById(ctx, problemId)
	return
}

func (mw InstrumentingMiddleware) FindAll(ctx context.Context) (output *string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "findall", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.FindAll(ctx)
	return
}

func (mw InstrumentingMiddleware) HealthCheck(ctx context.Context) (output *pb.HealthCheckResponse, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "healthcheck", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.HealthCheck(ctx)
	return
}
