package api

import (
	"encoding/json"
	"net/http"

	"github.com/dagmawiyohannes3914/task-queue-platform/internal/logger"
	"github.com/dagmawiyohannes3914/task-queue-platform/internal/models"
	"github.com/dagmawiyohannes3914/task-queue-platform/internal/repository"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SubmitJobRequest struct {
	Type    string                 `json:"type" validate:"required"`
	Payload map[string]interface{} `json:"payload" validate:"required"`
}

type SubmitJobResponse struct {
	ID string `json:"id"`
}

func SubmitJobHandler(w http.ResponseWriter, r *http.Request) {
	var req SubmitJobRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// Validate type
	if req.Type == "" || req.Payload == nil {
		http.Error(w, "type and payload are required", http.StatusBadRequest)
		return
	}

	payloadJSON, err := json.Marshal(req.Payload)
	if err != nil {
		http.Error(w, "failed to encode payload", http.StatusInternalServerError)
		return
	}

	job := models.Job{
		ID:         uuid.New(),
		Type:       req.Type,
		Payload:    payloadJSON,
		Status:     models.StatusPending,
		RetryCount: 0,
		MaxRetries: 3,
	}

	err = repository.DB.Create(&job).Error
	if err != nil {
		http.Error(w, "failed to create job", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(SubmitJobResponse{ID: job.ID.String()})
}

func GetJobStatusHandler(w http.ResponseWriter, r *http.Request) {
	jobIDParam := chi.URLParam(r, "id")
	jobID, err := uuid.Parse(jobIDParam)
	if err != nil {
		http.Error(w, "invalid job ID", http.StatusBadRequest)
		return
	}

	var job models.Job
	result := repository.DB.First(&job, "id = ?", jobID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, "job not found", http.StatusNotFound)
			return
		}
		logger.Log.Error("failed to fetch job", zap.Error(result.Error))
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(job)
}
