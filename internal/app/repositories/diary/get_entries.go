package diaryRepo

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/zura-t/observer.dev/internal/app/models"
	diaryUsecase "github.com/zura-t/observer.dev/internal/app/usecases/diary"
)

func (r *repo) GetEntries(ctx context.Context, filter *diaryUsecase.DiarySearchFilter) (*[]models.Diary, error) {
	qb := squirrel.Select(tableDiaryEntriesColumns...).
		From(tableDiaryEntries).
		Where(squirrel.Eq{columnUserID: filter.UserID}).
		OrderBy(columnEntryDate + " DESC").
		PlaceholderFormat(squirrel.Dollar)
	qb = applyDiarySearchFilter(qb, filter)

	sql, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("squirrel: %w", err)
	}

	rows, err := r.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("postgres: %w", convertPGError(err))
	}
	defer rows.Close()

	var entries []models.Diary
	for rows.Next() {
		var entry models.Diary
		if err := rows.Scan(
			&entry.ID,
			&entry.Title,
			&entry.Text,
			&entry.EntryDate,
			&entry.UserID,
			&entry.CreatedAt,
			&entry.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("postgres scan: %w", err)
		}
		entries = append(entries, entry)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("postgres rows: %w", err)
	}

	return &entries, nil
}

func applyDiarySearchFilter(qb squirrel.SelectBuilder, filter *diaryUsecase.DiarySearchFilter) squirrel.SelectBuilder {
	if filter == nil {
		return qb
	}

	if filter.DateFrom != nil {
		qb = qb.Where(squirrel.GtOrEq{columnEntryDate: *filter.DateFrom})
	}

	if filter.DateTo != nil {
		qb = qb.Where(squirrel.LtOrEq{columnEntryDate: *filter.DateTo})
	}

	return qb
}
