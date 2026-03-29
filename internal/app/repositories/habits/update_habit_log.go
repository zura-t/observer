package habitsRepo

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/zura-t/observer.dev/internal/app/models"
	habitsUsecase "github.com/zura-t/observer.dev/internal/app/usecases/habits"
)

func (r *repo) UpdateHabitLog(ctx context.Context, habitLog *habitsUsecase.UpdateHabitLog) (*models.HabitLog, error) {
	qb := squirrel.Update(tableHabitLogs).
		Where(squirrel.Eq{columnID: habitLog.ID}).
		Where(squirrel.Expr("habit_id IN (SELECT id FROM habits WHERE user_id = ?)", habitLog.UserID)).
		Suffix("RETURNING *").
		PlaceholderFormat(squirrel.Dollar)
	qb = applyHabitLogUpdate(qb, habitLog)

	sql, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("squirrel: %w", err)
	}

	var habitLogUpdated models.HabitLog
	err = r.db.QueryRow(ctx, sql, args...).Scan(
		&habitLogUpdated.ID,
		&habitLogUpdated.HabitID,
		&habitLogUpdated.LogDate,
		&habitLogUpdated.ActualCount,
		&habitLogUpdated.Note,
		&habitLogUpdated.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("postgres: %w", convertPGError(err, habitLogErrMap))
	}

	return &habitLogUpdated, nil
}

func applyHabitLogUpdate(qb squirrel.UpdateBuilder, habitLog *habitsUsecase.UpdateHabitLog) squirrel.UpdateBuilder {
	if !habitLog.LogDate.IsZero() {
		qb = qb.Set(columnLogDate, habitLog.LogDate)
	}

	if habitLog.ActualCount != nil {
		qb = qb.Set(columnActualCount, *habitLog.ActualCount)
	}

	if habitLog.Note != "" {
		qb = qb.Set(columnNote, habitLog.Note)
	}

	return qb
}
