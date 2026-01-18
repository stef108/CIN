package main

import (
	"html/template"
	"log"
	"net/http"

	"devops-project/config"
	"devops-project/handlers"
)

func main() {
	// Load Configuration
	cfg := config.LoadConfig()

	tpl := template.Must(template.ParseGlob("templates/*.html"))

	// Setup Routes
	http.HandleFunc("/", handlers.ListHandler(cfg, tpl))
	http.HandleFunc("/add", handlers.AddHandler)
	http.HandleFunc("/update", handlers.UpdateHandler)
	http.HandleFunc("/delete", handlers.DeleteHandler)
	http.HandleFunc("/health", handlers.HealthHandler)

	// Start Server
	log.Printf("Starting %s (v%s) on port %s...", cfg.AppName, cfg.Version, cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		log.Fatal(err)
	}
}
