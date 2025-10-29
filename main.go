package main

import (
	"log"
	"net/http"

	"github.com/irvaniamirali/go-prometheus-exporter/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// HTTP handler for health check
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	// Register custom collectors
	prometheus.MustRegister(metrics.NewSystemCollector())
	prometheus.MustRegister(metrics.NewAppCollector())

	// Expose /metrics endpoint
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
            <h1>Go Custom Exporter</h1>
            <p><a href="/metrics">Metrics</a> | <a href="/health">Health</a></p>
        `))
	})

	log.Println("Exporter is running on :9091")
	log.Fatal(http.ListenAndServe(":9091", nil))
}
