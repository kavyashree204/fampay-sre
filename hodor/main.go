package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "development"
	}

	http.HandleFunc("/hodr/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"service": "hodor",
			"message": "Hodor!",
			"env":     appEnv,
			"status":  "ok",
		})
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
	})

	fmt.Println("Hodor service running on port 8080")
	http.ListenAndServe(":8080", nil)
}