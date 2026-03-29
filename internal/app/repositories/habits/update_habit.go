package habitsRepo

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/zura-t/observer.dev/internal/app/models"
	habitsUsecase "github.com/zura-t/observer.dev/internal/app/usecases/habits"
)

func (r *repo) UpdateHabit(ctx context.Context, habit *habitsUsecase.UpdateHabit) (*models.Habit, error) {
	qb := squirrel.Update(tableHabits).
		Where(squirrel.Eq{columnID: habit.ID}).
		Where(squirrel.Eq{columnUserID: habit.UserID}).
		Suffix("RETURNING *").
		PlaceholderFormat(squirrel.Dollar)
	qb = applyHabitUpdate(qb, habit)

	sql, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("squirrel: %w", err)
	}

	var habitUpdated models.Habit
	err = r.db.QueryRow(ctx, sql, args...).Scan(
		&habitUpdated.ID,
		&habitUpdated.Title,
		&habitUpdated.Frequency,
		&habitUpdated.TargetCount,
		&habitUpdated.UserID,
		&habitUpdated.CreatedAt,
		&habitUpdated.UpdatedAt,
		&habitUpdated.DeletedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("postgres: %w", convertPGError(err, habitErrMap))
	}

	return &habitUpdated, nil
}

func applyHabitUpdate(qb squirrel.UpdateBuilder, habit *habitsUsecase.UpdateHabit) squirrel.UpdateBuilder {
	if habit.Title != "" {
		qb = qb.Set(columnTitle, habit.Title)
	}

	if habit.Frequency != "" {
		qb = qb.Set(columnFrequency, habit.Frequency)
	}

	if habit.TargetCount != nil {
		qb = qb.Set(columnTargetCount, *habit.TargetCount)
	}

	return qb
}
