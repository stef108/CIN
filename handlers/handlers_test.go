package handlers

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"devops-project/config"
)

func TestHealthHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Health check failed: got %v want %v", status, http.StatusOK)
	}
}

func TestListHandler(t *testing.T) {
	cfg := config.AppConfig{AppName: "TestApp", Version: "1.0"}

	dummyTpl := template.Must(template.New("index.html").Parse("<html>{{.AppName}}</html>"))

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := ListHandler(cfg, dummyTpl)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("List handler failed: got %v want %v", status, http.StatusOK)
	}
}
