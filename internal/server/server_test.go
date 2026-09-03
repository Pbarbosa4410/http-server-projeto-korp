package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
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

func TestProjetoKorpEndpoint(t *testing.T) {
	srv := New()

	req := httptest.NewRequest(http.MethodGet, "/projeto-korp", nil)
	rec := httptest.NewRecorder()

	srv.Handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status esperado %d, recebido %d", http.StatusOK, rec.Code)
	}

	contentType := rec.Header().Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		t.Errorf("Content-Type esperado application/json, recebido %q", contentType)
	}

	var response ProjetoKorpResponse

	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("erro ao decodificar resposta JSON: %v", err)
	}

	if response.Nome != "Projeto Korp" {
		t.Errorf(
			"nome esperado %q, recebido %q",
			"Projeto Korp",
			response.Nome,
		)
	}

	if response.Horario == "" {
		t.Fatal("campo horario não pode estar vazio")
	}

	parsedTime, err := time.Parse(time.RFC3339, response.Horario)
	if err != nil {
		t.Fatalf(
			"horario deveria estar no formato RFC3339/UTC, recebido %q: %v",
			response.Horario,
			err,
		)
	}

	if parsedTime.Location() != time.UTC {
		t.Errorf(
			"horario deveria estar em UTC, recebido %q",
			response.Horario,
		)
	}
}

func TestProjetoKorpSomenteGET(t *testing.T) {
	srv := New()

	req := httptest.NewRequest(http.MethodPost, "/projeto-korp", nil)
	rec := httptest.NewRecorder()

	srv.Handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf(
			"status esperado %d para POST, recebido %d",
			http.StatusMethodNotAllowed,
			rec.Code,
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
