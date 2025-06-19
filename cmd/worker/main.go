package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/dagmawiyohannes3914/task-queue-platform/internal/config"
	"github.com/dagmawiyohannes3914/task-queue-platform/internal/logger"
	"github.com/dagmawiyohannes3914/task-queue-platform/internal/models"
	"github.com/dagmawiyohannes3914/task-queue-platform/internal/queue"
	"github.com/dagmawiyohannes3914/task-queue-platform/internal/repository"
	"github.com/nats-io/nats.go"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func main() {
	config.LoadConfig()
	logger.InitLogger()
	repository.InitDB()
	queue.InitNATS(config.AppConfig.NatsURL)

	log.Println("Worker started and subscribed to jobs.new")

	_, err := queue.NatsConn.Subscribe("jobs.new", func(msg *nats.Msg) {
		handleJobMessage(msg)
	})
	if err != nil {
		log.Fatalf("Failed to subscribe to jobs.new: %v", err)
	}

	select {} // keep worker running forever
}

func handleJobMessage(msg *nats.Msg) {
	var payload map[string]string
	err := json.Unmarshal(msg.Data, &payload)
	if err != nil {
		log.Println("Failed to parse message:", err)
		return
	}

	jobIDStr := payload["job_id"]
	jobID, err := uuid.Parse(jobIDStr)
	if err != nil {
		log.Println("Invalid job ID:", err)
		return
	}

	processJob(jobID)
}

func processJob(jobID uuid.UUID) {
	var job models.Job
	result := repository.DB.First(&job, "id = ?", jobID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			log.Println("Job not found:", jobID)
		} else {
			log.Println("DB error:", result.Error)
		}
		return
	}

	// Update status to processing
	job.Status = models.StatusProcessing
	repository.DB.Save(&job)

	// Simulate work
	log.Printf("Processing job: %s (Type: %s)", job.ID, job.Type)
	time.Sleep(5 * time.Second) // simulate work

	// Simulate success or failure
	if simulateFailure() {
		job.Status = models.StatusFailed
		job.RetryCount++
		job.Error = "Simulated processing failure"

		if job.RetryCount < job.MaxRetries {
			job.Status = models.StatusRetrying

			// Re-publish to NATS for retry
			event := map[string]interface{}{"job_id": job.ID.String()}
			eventData, _ := json.Marshal(event)
			_ = queue.Publish("jobs.new", eventData)
		}

	} else {
		job.Status = models.StatusSuccess
		job.Error = ""
	}

	job.UpdatedAt = time.Now()
	repository.DB.Save(&job)
	log.Printf("Job %s updated to status: %s", job.ID, job.Status)
}

func simulateFailure() bool {
	// 20% chance to fail
	return time.Now().UnixNano()%5 == 0
}
