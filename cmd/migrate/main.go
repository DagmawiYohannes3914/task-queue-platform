package main

import (
	"github.com/dagmawiyohannes3914/task-queue-platform/internal/config"
	"github.com/dagmawiyohannes3914/task-queue-platform/internal/logger"
	"github.com/dagmawiyohannes3914/task-queue-platform/internal/repository"
)

func main() {
	config.LoadConfig()
	logger.InitLogger()
	repository.InitDB()
	repository.Migrate()
}
