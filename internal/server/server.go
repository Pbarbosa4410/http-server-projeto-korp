package server

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var requestsTotal = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total de requisições HTTP recebidas",
	},
)

func init() {
	prometheus.MustRegister(requestsTotal)
}

func New() *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestsTotal.Inc()
		fmt.Fprintln(w, "HTTP Server Projeto Korp funcionando!")
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		requestsTotal.Inc()
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	})

	mux.Handle("/metrics", promhttp.Handler())

	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}
