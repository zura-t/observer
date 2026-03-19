package diaryRepo

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/zura-t/observer.dev/internal/app/models"
	diaryUsecase "github.com/zura-t/observer.dev/internal/app/usecases/diary"
)

func (r *repo) UpdateDiaryEntry(ctx context.Context, entry *diaryUsecase.UpdateDiaryEntry) (*models.Diary, error) {
	qb := squirrel.Update(tableDiaryEntries).
		Where(squirrel.Eq{columnID: entry.ID}).
		Suffix("RETURNING *").
		PlaceholderFormat(squirrel.Dollar)
	qb = applyDiaryEntryUpdate(qb, entry)

	sql, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("squirrel: %w", err)
	}

	var diaryUpdated models.Diary
	err = r.db.QueryRow(ctx, sql, args...).Scan(
		&diaryUpdated.ID,
		&diaryUpdated.Title,
		&diaryUpdated.Text,
		&diaryUpdated.EntryDate,
		&diaryUpdated.UserID,
		&diaryUpdated.CreatedAt,
		&diaryUpdated.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("postgres: %w", convertPGError(err))
	}

	return &diaryUpdated, nil
}

func applyDiaryEntryUpdate(qb squirrel.UpdateBuilder, diaryEntry *diaryUsecase.UpdateDiaryEntry) squirrel.UpdateBuilder {
	if diaryEntry.Title != "" {
		qb = qb.Set(columnTitle, diaryEntry.Title)
	}

	if diaryEntry.Text != "" {
		qb = qb.Set(columnText, diaryEntry.Text)
	}

	if diaryEntry.EntryDate != nil {
		qb = qb.Set(columnEntryDate, *diaryEntry.EntryDate)
	}

	qb = qb.Set(columnUpdatedAt, squirrel.Expr("NOW()"))

	return qb
}
