package problm

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/metrics"
	"github.com/jalgoarena/problems/pb"
	"time"
)

type instrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	Next           ProblemsService
}

func InstrumentingMiddleware(requestCount metrics.Counter, requestLatency metrics.Histogram) ServiceMiddleware {
	return func(next ProblemsService) ProblemsService {
		return instrumentingMiddleware{requestCount, requestLatency, next}
	}
}

func (mw instrumentingMiddleware) FindById(ctx context.Context, problemId string) (output *pb.Problem, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "findbyid", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.FindById(ctx, problemId)
	return
}

func (mw instrumentingMiddleware) FindAll(ctx context.Context) (output *string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "findall", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.FindAll(ctx)
	return
}

func (mw instrumentingMiddleware) HealthCheck(ctx context.Context) (output *pb.HealthCheckResponse, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "healthcheck", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.HealthCheck(ctx)
	return
}
