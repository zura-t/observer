package habitsRepo

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/zura-t/observer.dev/internal/app/models"
)

func (r *repo) GetHabits(ctx context.Context, userID uint64) (*[]models.Habit, error) {
	qb := squirrel.Select(tableHabitsColumns...).
		From(tableHabits).
		Where(squirrel.Eq{"user_id": userID}).
		OrderBy(columnTitle + " ASC").
		PlaceholderFormat(squirrel.Dollar)

	sql, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("squirrel: %w", err)
	}

	rows, err := r.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("postgres: %w", convertPGError(err, habitErrMap))
	}
	defer rows.Close()

	var habits []models.Habit
	for rows.Next() {
		var habit models.Habit
		if err := rows.Scan(
			&habit.ID,
			&habit.Title,
			&habit.Frequency,
			&habit.TargetCount,
			&habit.UserID,
			&habit.CreatedAt,
			&habit.UpdatedAt,
			&habit.DeletedAt,
		); err != nil {
			return nil, fmt.Errorf("postgres scan: %w", err)
		}
		habits = append(habits, habit)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("postgres rows: %w", err)
	}

	return &habits, nil
}
