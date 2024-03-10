package main

/*Logging and Monitoring: Centralized logging and monitoring of
each individual service is essential to keep tabs on system health and debug
 issues in a microservices architecture.
*/
import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "app_requests_total",
		Help: "Total number of requests",
	})

	requestDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "app_request_duration_seconds",
		Help:    "Request duration in seconds",
		Buckets: prometheus.LinearBuckets(0.01, 0.01, 10),
	})
)

func init() {
	prometheus.MustRegister(requestsTotal)
	prometheus.MustRegister(requestDuration)
}

func main() {
	http.HandleFunc("/", handleRequest)
	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	// Simulating some work
	time.Sleep(100 * time.Millisecond)

	w.Write([]byte("Hello, World!"))

	duration := float64(time.Since(startTime)) / float64(time.Second)
	requestsTotal.Inc()
	requestDuration.Observe(duration)
}
