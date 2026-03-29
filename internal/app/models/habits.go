package models

import "time"

type Habit struct {
	ID          uint64
	Title       string
	Frequency   Frequency
	TargetCount uint8
	UserID      uint64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

type Frequency string

const (
	FrequencyDaily  Frequency = "daily"
	FrequencyWeekly Frequency = "weekly"
)

type HabitLog struct {
	ID          uint64
	HabitID     uint64
	LogDate     time.Time
	ActualCount uint8
	Note        string
	CreatedAt   time.Time
}
