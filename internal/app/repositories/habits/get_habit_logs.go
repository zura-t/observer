package habitsRepo

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/zura-t/observer.dev/internal/app/models"
	habitsUsecase "github.com/zura-t/observer.dev/internal/app/usecases/habits"
)

func (r *repo) GetHabitLogs(ctx context.Context, filter *habitsUsecase.GetHabitLogsFilter) (*[]models.HabitLog, error) {
	qb := squirrel.Select(tableHabitLogsColumns...).
		From(tableHabitLogs).
		Where(squirrel.Eq{columnHabitID: filter.HabitID}).
		Where(squirrel.Expr("habit_id IN (SELECT id FROM habits WHERE user_id = ?)", filter.UserID)).
		OrderBy(columnCreatedAt + " DESC").
		Limit(uint64(filter.Limit)).
		Offset(uint64(filter.Offset)).
		PlaceholderFormat(squirrel.Dollar)

	sql, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("squirrel: %w", err)
	}

	rows, err := r.db.Query(ctx, sql, args...)
	if err != nil {
		r.logger.Error(err)
		return nil, fmt.Errorf("postgres: %w", convertPGError(err, habitLogErrMap))
	}
	defer rows.Close()

	var habitLogs []models.HabitLog
	for rows.Next() {
		var habitLog models.HabitLog
		if err := rows.Scan(
			&habitLog.ID,
			&habitLog.HabitID,
			&habitLog.LogDate,
			&habitLog.ActualCount,
			&habitLog.Note,
			&habitLog.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("postgres scan: %w", err)
		}
		habitLogs = append(habitLogs, habitLog)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("postgres rows: %w", err)
	}

	return &habitLogs, nil
}
