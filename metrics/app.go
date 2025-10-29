package metrics

import (
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// AppCollector tracks application-level metrics
type AppCollector struct {
	requestsTotal   *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
	mu              sync.Mutex
}

var appCollector *AppCollector

func NewAppCollector() *AppCollector {
	ac := &AppCollector{
		requestsTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_requests_total",
				Help: "Total number of HTTP requests",
			},
			[]string{"path", "method"},
		),
		requestDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "http_request_duration_seconds",
				Help:    "Duration of HTTP requests in seconds",
				Buckets: prometheus.DefBuckets,
			},
			[]string{"path", "method"},
		),
	}

	// Simulate real app traffic
	go ac.simulateTraffic()
	appCollector = ac
	return ac
}

func (ac *AppCollector) Describe(ch chan<- *prometheus.Desc) {
	ac.requestsTotal.Describe(ch)
	ac.requestDuration.Describe(ch)
}

func (ac *AppCollector) Collect(ch chan<- prometheus.Metric) {
	ac.requestsTotal.Collect(ch)
	ac.requestDuration.Collect(ch)
}

// simulateTraffic generates fake traffic for demo
func (ac *AppCollector) simulateTraffic() {
	ticker := time.NewTicker(2 * time.Second)
	paths := []string{"/api/users", "/api/orders", "/health"}
	methods := []string{"GET", "POST"}

	for range ticker.C {
		path := paths[time.Now().Unix()%int64(len(paths))]
		method := methods[time.Now().Unix()%int64(len(methods))]
		duration := 0.1 + (float64(time.Now().UnixNano()%1000) / 1000.0) // 0.1 to 0.2s

		ac.requestsTotal.WithLabelValues(path, method).Inc()
		ac.requestDuration.WithLabelValues(path, method).Observe(duration)
	}
}

// RecordRequest records real HTTP requests (optional middleware)
func RecordRequest(path, method string, duration float64) {
	if appCollector != nil {
		appCollector.mu.Lock()
		defer appCollector.mu.Unlock()
		appCollector.requestsTotal.WithLabelValues(path, method).Inc()
		appCollector.requestDuration.WithLabelValues(path, method).Observe(duration)
	}
}
