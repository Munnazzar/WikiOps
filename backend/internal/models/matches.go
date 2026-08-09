package models

import (
	"time"

	"github.com/google/uuid"
)

type Matches struct {
	ID              uuid.UUID `gorm:"type:uuid;default;gen_random_uuid();primaryKey"`
	LinkID          string    `gorm:"type:varchar(7);not null"`
	StartPage       string    `gorm:"not null"`
	TargetPage      string    `gorm:"not null"`
	WinnerID        string    `gorm:"type:varchar(255);not null"`
	LoserID         string    `gorm:"type:varchar(255);not null"`
	DurationSeconds int       `gorm:"not null"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
}
