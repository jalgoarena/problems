package problm

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func decodeProblemRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	problemId, ok := vars["problemId"]
	if !ok {
		return nil, ErrBadRouting
	}

	return problemRequest{
		problemId,
	}, nil
}

func encodeProblemResponse(_ context.Context, w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(resp.(problemResponse).Problem)
}

func decodeProblemsRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return problemsRequest{}, nil
}

func encodeProblemsResponse(_ context.Context, w http.ResponseWriter, resp interface{}) error {
	tmp := resp.(problemsResponse).Problems
	problemsJSON := *tmp
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err := w.Write([]byte(problemsJSON))
	return err
}

func decodeHealthCheckRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return healthCheckRequest{}, nil
}

func encodeHealthCheckResponse(_ context.Context, w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(resp.(healthCheckResponse).HealthCheckResult)
}

func MakeHTTPHandler(endpoints Endpoints, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	r.Methods("GET").Path("/api/v1/problems").Handler(httptransport.NewServer(
		endpoints.ProblemsEndpoint,
		decodeProblemsRequest,
		encodeProblemsResponse,
	))
	r.Methods("GET").Path("/api/v1/problems/{problemId}").Handler(httptransport.NewServer(
		endpoints.ProblemEndpoint,
		decodeProblemRequest,
		encodeProblemResponse,
	))
	r.Methods("GET").Path("/metrics").Handler(promhttp.Handler())
	r.Methods("GET").Path("/health").Handler(httptransport.NewServer(
		endpoints.HealthCheckEndpoint,
		decodeHealthCheckRequest,
		encodeHealthCheckResponse,
	))

	return httpLoggingMiddleware(logger)(r)
}

type httpMiddleware func(http.Handler) http.Handler

func httpLoggingMiddleware(logger log.Logger) httpMiddleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Log("transport", "http", "before", r.RequestURI)
			defer logger.Log("transport", "http", "after", r.RequestURI)
			h.ServeHTTP(w, r)
		})
	}
}
