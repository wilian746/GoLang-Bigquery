package api

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/GoLang-Bigquery/handlers"
)

func Server() *http.ServeMux {
	mu := http.NewServeMux()
	mu.Handle("/metrics", promhttp.Handler())
	mu.HandleFunc("/api", handlers.NewHelloWorldHandler().HelloWorldHandler)
	mu.HandleFunc("/production/", handlers.NewProductionHandler().ProductionHandler)
	return mu
}
