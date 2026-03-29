package habitsRepo

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/zura-t/observer.dev/internal/app/models"
)

func (r *repo) GetHabit(ctx context.Context, id uint64, userID uint64) (*models.Habit, error) {
	sql, args, err := squirrel.Select(tableHabitsColumns...).
		From(tableHabits).
		Where(squirrel.Eq{columnID: id, columnUserID: userID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("squirrel: %w", err)
	}

	var habit models.Habit
	err = r.db.QueryRow(ctx, sql, args...).Scan(
		&habit.ID,
		&habit.Title,
		&habit.Frequency,
		&habit.TargetCount,
		&habit.UserID,
		&habit.CreatedAt,
		&habit.UpdatedAt,
		&habit.DeletedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("postgres: %w", convertPGError(err, habitErrMap))
	}

	return &habit, nil
}
