package habitsRepo

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/zura-t/observer.dev/internal/app/models"
	habitsUsecase "github.com/zura-t/observer.dev/internal/app/usecases/habits"
)

func (r *repo) CreateHabit(ctx context.Context, habit *habitsUsecase.CreateHabit) (*models.Habit, error) {
	newEntry := newHabit(habit)
	sql, args, err := squirrel.Insert(tableHabits).
		Columns(createHabitColumns...).
		Values(newEntry.Values(createHabitColumns...)...).
		Suffix("Returning *").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("squirrel: %w", err)
	}

	var habitCreated models.Habit
	err = r.db.QueryRow(ctx, sql, args...).Scan(
		&habitCreated.ID,
		&habitCreated.Title,
		&habitCreated.Frequency,
		&habitCreated.TargetCount,
		&habitCreated.UserID,
		&habitCreated.CreatedAt,
		&habitCreated.UpdatedAt,
		&habitCreated.DeletedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("postgres: %w", convertPGError(err, habitErrMap))
	}

	return &habitCreated, nil
}
