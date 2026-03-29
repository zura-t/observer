package habitsRepo

import (
	"strings"
	"time"
)

const tableHabits = "public.habits"
const tableHabitLogs = "public.habit_logs"

const (
	columnID          = "id"
	columnTitle       = "title"
	columnFrequency   = "frequency"
	columnTargetCount = "target_count"
	columnUserID      = "user_id"
	columnCreatedAt   = "created_at"
	columnUpdatedAt   = "updated_at"
	columnDeletedAt   = "deleted_at"

	columnHabitID           = "habit_id"
	columnLogDate           = "log_date"
	columnActualCount       = "actual_count"
	columnNote              = "note"
	columnHabitLogID        = tableHabitLogs + "." + columnID
	columnHabitLogCreatedAt = tableHabitLogs + "." + columnCreatedAt
)

type habit struct {
	Title       string    `db:"title"`
	Frequency   string    `db:"frequency"`
	TargetCount uint8     `db:"target_count"`
	UserID      uint64    `db:"user_id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type habitLog struct {
	HabitID     uint64    `db:"habit_id"`
	LogDate     time.Time `db:"log_date"`
	ActualCount uint8     `db:"actual_count"`
	Note        string    `db:"note"`
	CreatedAt   time.Time `db:"created_at"`
}

var (
	tableHabitsColumns = []string{
		columnID,
		columnTitle,
		columnFrequency,
		columnTargetCount,
		columnUserID,
		columnCreatedAt,
		columnUpdatedAt,
		columnDeletedAt,
	}

	createHabitColumns = []string{
		columnTitle,
		columnFrequency,
		columnTargetCount,
		columnUserID,
	}

	tableHabitLogsColumns = []string{
		columnHabitLogID,
		columnHabitID,
		columnLogDate,
		columnActualCount,
		columnNote,
		columnHabitLogCreatedAt,
	}

	createHabitLogColumns = []string{
		columnHabitID,
		columnLogDate,
		columnActualCount,
		columnNote,
	}
)

func (o *habit) mapFields() map[string]any {
	return map[string]any{
		columnTitle:       o.Title,
		columnFrequency:   o.Frequency,
		columnTargetCount: o.TargetCount,
		columnUserID:      o.UserID,
		columnCreatedAt:   o.CreatedAt,
		columnUpdatedAt:   o.UpdatedAt,
	}
}

func (o *habit) Values(columns ...string) []any {
	mapFields := o.mapFields()
	values := make([]any, 0, len(columns))
	for i := range columns {
		if v, ok := mapFields[columns[i]]; ok {
			values = append(values, v)
		} else {
			values = append(values, nil)
		}
	}
	return values
}

func (o *habit) ReturningValues(columns ...string) string {
	return strings.Join(columns, ", ")
}

func (o *habitLog) mapFields() map[string]any {
	return map[string]any{
		columnHabitID:     o.HabitID,
		columnLogDate:     o.LogDate,
		columnActualCount: o.ActualCount,
		columnNote:        o.Note,
		columnCreatedAt:   o.CreatedAt,
	}
}

func (o *habitLog) Values(columns ...string) []any {
	mapFields := o.mapFields()
	values := make([]any, 0, len(columns))
	for i := range columns {
		if v, ok := mapFields[columns[i]]; ok {
			values = append(values, v)
		} else {
			values = append(values, nil)
		}
	}
	return values
}

func (o *habitLog) ReturningValues(columns ...string) string {
	return strings.Join(columns, ", ")
}
