package habitsController

import "time"

type CreateHabitRequest struct {
	Title       string `json:"title" binding:"required"`
	Frequency   string `json:"frequency" binding:"required,oneof=daily weekly"`
	TargetCount uint8  `json:"target_count" binding:"required,min=1"`
}

type GetHabitByIDRequest struct {
	ID uint64 `uri:"id" binding:"required"`
}

type UpdateHabitRequest struct {
	Title       string `json:"title"`
	Frequency   string `json:"frequency" binding:"omitempty,oneof=daily weekly"`
	TargetCount *uint8 `json:"target_count" binding:"omitempty,min=1"`
}

type DeleteHabitRequest struct {
	ID uint64 `uri:"id" binding:"required"`
}

type HabitIDRequest struct {
	HabitID uint64 `uri:"habitID" binding:"required"`
}

type CreateHabitLogRequest struct {
	LogDate     time.Time `json:"log_date" binding:"required"`
	ActualCount uint8     `json:"actual_count" binding:"required,min=1"`
	Note        string    `json:"note"`
}

type GetHabitLogByIDRequest struct {
	ID uint64 `uri:"id" binding:"required"`
}

type GetHabitLogsRequest struct {
	Limit  uint8 `form:"limit,default=10"`
	Offset uint8 `form:"offset,default=0"`
}

type UpdateHabitLogRequest struct {
	LogDate     time.Time `json:"log_date"`
	ActualCount *uint8    `json:"actual_count" binding:"omitempty,min=1"`
	Note        string    `json:"note"`
}

type DeleteHabitLogRequest struct {
	ID uint64 `uri:"id" binding:"required"`
}
