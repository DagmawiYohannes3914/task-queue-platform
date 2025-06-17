package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type JobStatus string

const (
	StatusPending    JobStatus = "pending"
	StatusProcessing JobStatus = "processing"
	StatusSuccess    JobStatus = "success"
	StatusFailed     JobStatus = "failed"
	StatusRetrying   JobStatus = "retrying"
)

type Job struct {
	ID         uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Type       string         `gorm:"size:50;not null"`
	Payload    datatypes.JSON `gorm:"type:jsonb;not null"`
	Status     JobStatus      `gorm:"size:20;not null"`
	RetryCount int            `gorm:"default:0"`
	MaxRetries int            `gorm:"default:3"`
	Error      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
