package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type JobStatus string

const (
	StatusPending    JobStatus = "pending"
	StatusProcessing JobStatus = "processing"
	StatusSuccess    JobStatus = "success"
	StatusFailed     JobStatus = "failed"
	StatusRetrying   JobStatus = "retrying"
)

// type Job struct {
// 	ID         uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
// 	Type       string         `gorm:"size:50;not null"`
// 	Payload    datatypes.JSON `gorm:"type:jsonb;not null"`
// 	Status     JobStatus      `gorm:"size:20;not null"`
// 	RetryCount int            `gorm:"default:0"`
// 	MaxRetries int            `gorm:"default:3"`
// 	Error      string
// 	CreatedAt  time.Time
// 	UpdatedAt  time.Time
// }
type Job struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID     uuid.UUID `gorm:"type:uuid;index" json:"user_id"`
	Type       string    `gorm:"not null" json:"type"`
	Payload    json.RawMessage `gorm:"type:jsonb" json:"payload"`
	Status     JobStatus  `gorm:"type:job_status" json:"status"`
	// Status     JobStatus       `gorm:"type:text;not null;default:'pending'" json:"status"`
	RetryCount int        `gorm:"default:0" json:"retry_count"`
	MaxRetries int        `gorm:"default:3" json:"max_retries"`
	Error      string     `gorm:"" json:"error"`
	CreatedAt  time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}

