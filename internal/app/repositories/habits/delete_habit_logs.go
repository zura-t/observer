package habitsRepo

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
)

func (r *repo) DeleteHabitLogs(ctx context.Context, habitID uint64, userID uint64) error {
	qb := squirrel.Delete(tableHabitLogs).
		Where(squirrel.Eq{columnHabitID: habitID}).
		Where(squirrel.Expr("habit_id IN (SELECT id FROM habits WHERE user_id = ?)", userID)).
		PlaceholderFormat(squirrel.Dollar)

	sql, args, err := qb.ToSql()
	if err != nil {
		return fmt.Errorf("squirrel: %w", err)
	}

	_, err = r.db.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("postgres: %w", convertPGError(err, habitLogErrMap))
	}

	return nil
}
