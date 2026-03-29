package habitsUsecase

import "time"

type CreateHabit struct {
	ID          uint64
	Title       string
	Frequency   string
	TargetCount uint8
	UserID      uint64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UpdateHabit struct {
	ID          uint64
	Title       string
	Frequency   string
	TargetCount *uint8
	UserID      uint64
}

type CreateHabitLog struct {
	ID          uint64
	HabitID     uint64
	LogDate     time.Time
	ActualCount uint8
	Note        string
	CreatedAt   time.Time
}

type UpdateHabitLog struct {
	ID          uint64
	HabitID     uint64
	LogDate     time.Time
	ActualCount *uint8
	Note        string
	UserID      uint64
}

type GetHabitLogsFilter struct {
	HabitID uint64
	UserID  uint64
	Limit   uint8
	Offset  uint8
}
