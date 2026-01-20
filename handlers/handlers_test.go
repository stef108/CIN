package handlers

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"devops-project/config"
	"devops-project/models"
)

// Helper to reset the global state for testing
func resetTasks() {
	mu.Lock()
	defer mu.Unlock()
	tasks = []models.Task{}
	nextID = 1
}

func TestAddHandler(t *testing.T) {
	resetTasks()

	// 1. Test POST request adds a task
	form := url.Values{}
	form.Add("title", "Integration Test Task")
	req, err := http.NewRequest("POST", "/add", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddHandler)
	handler.ServeHTTP(rr, req)

	// Check Redirect
	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("AddHandler returned wrong status code: got %v want %v",
			status, http.StatusSeeOther)
	}

	// Check if task was actually added to the global slice
	mu.Lock()
	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	} else if tasks[0].Title != "Integration Test Task" {
		t.Errorf("Expected task title 'Integration Test Task', got '%s'", tasks[0].Title)
	}
	mu.Unlock()

	// 2. Test GET request redirects (Method Not Allowed logic)
	reqGet, _ := http.NewRequest("GET", "/add", nil)
	rrGet := httptest.NewRecorder()
	handler.ServeHTTP(rrGet, reqGet)

	if status := rrGet.Code; status != http.StatusSeeOther {
		t.Errorf("AddHandler GET returned wrong status code: got %v want %v",
			status, http.StatusSeeOther)
	}
}

func TestUpdateHandler(t *testing.T) {
	resetTasks()

	// Pre-seed a task
	mu.Lock()
	tasks = append(tasks, models.Task{ID: 1, Title: "Task to Update", Status: models.StatusTodo})
	mu.Unlock()

	// Create form to update ID 1
	form := url.Values{}
	form.Add("id", "1")
	req, err := http.NewRequest("POST", "/update", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("UpdateHandler returned wrong status code: got %v want %v",
			status, http.StatusSeeOther)
	}

	// Verify status flipped from TODO to DONE
	mu.Lock()
	if tasks[0].Status != models.StatusDone {
		t.Errorf("Expected status DONE, got %s", tasks[0].Status)
	}
	mu.Unlock()
}

func TestDeleteHandler(t *testing.T) {
	resetTasks()

	// Pre-seed a task
	mu.Lock()
	tasks = append(tasks, models.Task{ID: 1, Title: "Task to Delete", Status: models.StatusTodo})
	mu.Unlock()

	form := url.Values{}
	form.Add("id", "1")
	req, err := http.NewRequest("POST", "/delete", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("DeleteHandler returned wrong status code: got %v want %v",
			status, http.StatusSeeOther)
	}

	// Verify task is gone
	mu.Lock()
	if len(tasks) != 0 {
		t.Errorf("Expected 0 tasks, got %d", len(tasks))
	}
	mu.Unlock()
}

func TestLoginPageHandler(t *testing.T) {
	cfg := config.AppConfig{AppName: "TestApp"}
	// Create a dummy login.html template for testing
	dummyTpl := template.Must(template.New("login.html").Parse("<html>{{.AppName}} - {{.Error}}</html>"))

	req, err := http.NewRequest("GET", "/login", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := LoginPageHandler(cfg, dummyTpl)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("LoginPageHandler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestLoginHandler(t *testing.T) {
	cfg := config.AppConfig{AppName: "TestApp"}
	dummyTpl := template.Must(template.New("login.html").Parse("<html>{{.Error}}</html>"))

	// 1. Test Successful Login
	form := url.Values{}
	form.Add("username", "admin")
	form.Add("password", "secret")
	req, _ := http.NewRequest("POST", "/auth/login", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := LoginHandler(cfg, dummyTpl)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("LoginHandler Success returned wrong status code: got %v want %v",
			status, http.StatusSeeOther)
	}

	// 2. Test Failed Login (Empty credentials)
	formEmpty := url.Values{}
	formEmpty.Add("username", "")
	formEmpty.Add("password", "")
	reqFail, _ := http.NewRequest("POST", "/auth/login", strings.NewReader(formEmpty.Encode()))
	reqFail.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rrFail := httptest.NewRecorder()
	handler.ServeHTTP(rrFail, reqFail)

	if status := rrFail.Code; status != http.StatusUnauthorized {
		t.Errorf("LoginHandler Fail returned wrong status code: got %v want %v",
			status, http.StatusUnauthorized)
	}
}
