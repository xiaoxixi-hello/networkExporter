package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"ylinyang.com/networkExporter/collector"
)

func main() {
	//flag.String("f", "", "-f /config/config.ini")
	//
	//flag.Parsed()

	var metricsNamespace = "network"

	metrics := collector.NewMetrics(metricsNamespace)
	registry := prometheus.NewRegistry()
	registry.MustRegister(metrics)

	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	log.Printf("Starting Server at http://localhost:%s%s", "9880", "/metrics")
	log.Fatal(http.ListenAndServe(":9880", nil))

}
