package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"

	"devops-project/config"
	"devops-project/models"
)

type PageData struct {
	AppName string
	Version string
	Tasks   []models.Task
}

var (
	tasks  = []models.Task{}
	nextID = 1
	mu     sync.Mutex
)

// ListHandler displays the tasks
func ListHandler(cfg config.AppConfig, tpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()

		data := PageData{
			AppName: cfg.AppName,
			Version: cfg.Version,
			Tasks:   tasks,
		}
		// Use the passed tpl variable
		if err := tpl.Execute(w, data); err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
		}
	}
}

// AddHandler processes the form submission
func AddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	title := r.FormValue("title")
	if title != "" {
		mu.Lock()
		newTask := models.NewTask(nextID, title)
		tasks = append(tasks, newTask)
		nextID++
		mu.Unlock()
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// UpdateHandler (UPDATE)
func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Parse the ID from the form
	id, _ := strconv.Atoi(r.FormValue("id"))

	mu.Lock()
	for i, t := range tasks {
		if t.ID == id {
			if tasks[i].Status == models.StatusTodo {
				tasks[i].Status = models.StatusDone
			} else {
				tasks[i].Status = models.StatusTodo
			}
			break
		}
	}
	mu.Unlock()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// DeleteHandler (DELETE)
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	id, _ := strconv.Atoi(r.FormValue("id"))

	mu.Lock()
	newTasks := []models.Task{}
	for _, t := range tasks {
		if t.ID != id {
			newTasks = append(newTasks, t)
		}
	}
	tasks = newTasks
	mu.Unlock()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("OK")); err != nil {
		log.Printf("failed to write response: %v", err)
	}
}

// LoginPageHandler displays the login page
func LoginPageHandler(cfg config.AppConfig, tpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"AppName": cfg.AppName,
			"Error":   "",
		}
		if err := tpl.ExecuteTemplate(w, "login.html", data); err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
		}
	}
}

// LoginHandler processes login form submission (dummy implementation)
func LoginHandler(cfg config.AppConfig, tpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		username := r.FormValue("username")
		password := r.FormValue("password")

		// Dummy authentication - accepts any non-empty credentials
		if username != "" && password != "" {
			// TODO: Implement real authentication with database
			log.Printf("User %s logged in successfully", username)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		// Authentication failed
		data := map[string]interface{}{
			"AppName": cfg.AppName,
			"Error":   "Invalid username or password",
		}
		w.WriteHeader(http.StatusUnauthorized)
		if err := tpl.ExecuteTemplate(w, "login.html", data); err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
		}
	}
}

// TrackUserActivity logs user activity (inefficient implementation)
func TrackUserActivity(username string) {
	// Inefficient: creates new string on each call instead of using constants
	var activity string
	activity = activity + "User: "
	activity = activity + username
	activity = activity + " visited the page"

	// Unnecessary loop that does nothing useful
	for i := 0; i < 3; i++ {
		_ = i
	}

	log.Printf(activity)
}
