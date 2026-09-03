package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRootEndpoint(t *testing.T) {
	srv := New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	srv.Handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status esperado %d, recebido %d", http.StatusOK, rec.Code)
	}

	expected := "HTTP Server Projeto Korp funcionando!"

	if strings.TrimSpace(rec.Body.String()) != expected {
		t.Errorf(
			"resposta esperada %q, recebida %q",
			expected,
			strings.TrimSpace(rec.Body.String()),
		)
	}
}

func TestHealthEndpoint(t *testing.T) {
	srv := New()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	srv.Handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status esperado %d, recebido %d", http.StatusOK, rec.Code)
	}

	expected := "OK"

	if strings.TrimSpace(rec.Body.String()) != expected {
		t.Errorf(
			"resposta esperada %q, recebida %q",
			expected,
			strings.TrimSpace(rec.Body.String()),
		)
	}
}

func TestMetricsEndpoint(t *testing.T) {
	srv := New()

	req := httptest.NewRequest(http.MethodGet, "/metrics", nil)
	rec := httptest.NewRecorder()

	srv.Handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status esperado %d, recebido %d", http.StatusOK, rec.Code)
	}

	body := rec.Body.String()

	if !strings.Contains(body, "http_requests_total") {
		t.Errorf("métrica http_requests_total não encontrada no endpoint /metrics")
	}
}
