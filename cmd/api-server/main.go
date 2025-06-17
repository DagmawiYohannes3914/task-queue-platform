package main

import (
	"fmt"
	"net/http"

	"github.com/dagmawiyohannes3914/task-queue-platform/internal/config"
	"github.com/dagmawiyohannes3914/task-queue-platform/internal/logger"
	"github.com/dagmawiyohannes3914/task-queue-platform/internal/repository"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func main() {
	config.LoadConfig()
	logger.InitLogger()
	repository.InitDB()
	
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API Server is running"))
	})

	port := config.AppConfig.ServerPort
	logger.Log.Info("Starting API server", zap.String("port", port))
	fmt.Println("API Server is running")
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		logger.Log.Fatal("Server failed", zap.Error(err))
	}
}
