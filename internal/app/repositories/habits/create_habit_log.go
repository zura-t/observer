package habitsRepo

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/zura-t/observer.dev/internal/app/models"
	habitsUsecase "github.com/zura-t/observer.dev/internal/app/usecases/habits"
)

func (r *repo) CreateHabitLog(ctx context.Context, habitLog *habitsUsecase.CreateHabitLog) (*models.HabitLog, error) {
	newEntry := newHabitLog(habitLog)
	sql, args, err := squirrel.Insert(tableHabitLogs).
		Columns(createHabitLogColumns...).
		Values(newEntry.Values(createHabitLogColumns...)...).
		Suffix("Returning *").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("squirrel: %w", err)
	}

	var habitLogCreated models.HabitLog
	err = r.db.QueryRow(ctx, sql, args...).Scan(
		&habitLogCreated.ID,
		&habitLogCreated.HabitID,
		&habitLogCreated.LogDate,
		&habitLogCreated.ActualCount,
		&habitLogCreated.Note,
		&habitLogCreated.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("postgres: %w", convertPGError(err, habitLogErrMap))
	}

	return &habitLogCreated, nil
}
