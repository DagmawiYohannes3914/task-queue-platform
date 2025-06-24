package repository

import (
	"log"

	"github.com/dagmawiyohannes3914/task-queue-platform/internal/models"
)

func Migrate() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Job{},
	)

	if err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}

	log.Println("Database migration completed successfully.")
}
