package habitsRepo

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
)

func (r *repo) DeleteHabit(ctx context.Context, id uint64, userID uint64) error {
	qb := squirrel.Delete(tableHabits).
		Where(squirrel.Eq{columnID: id, columnUserID: userID}).
		PlaceholderFormat(squirrel.Dollar)

	sql, args, err := qb.ToSql()
	if err != nil {
		return fmt.Errorf("squirrel: %w", err)
	}

	_, err = r.db.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("postgres: %w", convertPGError(err, habitErrMap))
	}

	return nil
}
