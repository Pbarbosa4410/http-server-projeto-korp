package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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

type ProjetoKorpResponse struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

func New() *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestsTotal.Inc()

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintln(w, "HTTP Server Projeto Korp funcionando!")
	})

	mux.HandleFunc("/projeto-korp", func(w http.ResponseWriter, r *http.Request) {
		requestsTotal.Inc()

		if r.Method != http.MethodGet {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}

		response := ProjetoKorpResponse{
			Nome:    "Projeto Korp",
			Horario: time.Now().UTC().Format(time.RFC3339),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Erro ao gerar resposta", http.StatusInternalServerError)
			return
		}
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
