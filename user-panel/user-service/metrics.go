
package main

import (
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "net/http"
    "time"
)

// Define Prometheus metrics
var (
    httpRequestsTotal = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests",
        },
        []string{"service", "endpoint", "method", "status"},
    )
    httpRequestDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "http_request_duration_seconds",
            Help:    "Histogram of response time for HTTP requests",
            Buckets: prometheus.DefBuckets,
        },
        []string{"service", "endpoint", "method"},
    )
)

func init() {
    prometheus.MustRegister(httpRequestsTotal)
    prometheus.MustRegister(httpRequestDuration)
}

// MetricsMiddleware to track metrics for each request
func MetricsMiddleware(service string, next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        endpoint := r.URL.Path
        method := r.Method

        // Start timer
        start := time.Now()
        rr := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}
        next.ServeHTTP(rr, r)

        // Observe metrics
        duration := time.Since(start).Seconds()
        status := rr.statusCode

        httpRequestsTotal.WithLabelValues(service, endpoint, method, http.StatusText(status)).Inc()
        httpRequestDuration.WithLabelValues(service, endpoint, method).Observe(duration)
    })
}

// responseRecorder to capture HTTP status code
type responseRecorder struct {
    http.ResponseWriter
    statusCode int
}

func (rr *responseRecorder) WriteHeader(statusCode int) {
    rr.statusCode = statusCode
    rr.ResponseWriter.WriteHeader(statusCode)
}

// MetricsHandler for Prometheus endpoint
func MetricsHandler() http.Handler {
    return promhttp.Handler()
}
