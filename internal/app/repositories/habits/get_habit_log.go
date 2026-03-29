package habitsRepo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/zura-t/observer.dev/internal/app/models"
)

func (r *repo) GetHabitLog(ctx context.Context, id uint64, userID uint64) (*models.HabitLog, error) {
	log.Print(time.Now().UTC())
	sql, args, err := squirrel.Select(tableHabitLogsColumns...).
		From(tableHabitLogs).
		Where(squirrel.Eq{columnID: id}).
		Where(squirrel.Expr("habit_id IN (SELECT id FROM habits WHERE user_id = ?)", userID)).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("squirrel: %w", err)
	}

	var habitLog models.HabitLog
	err = r.db.QueryRow(ctx, sql, args...).Scan(
		&habitLog.ID,
		&habitLog.HabitID,
		&habitLog.LogDate,
		&habitLog.ActualCount,
		&habitLog.Note,
		&habitLog.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("postgres: %w", convertPGError(err, habitLogErrMap))
	}

	return &habitLog, nil
}
