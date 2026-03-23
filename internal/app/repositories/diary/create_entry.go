package diaryRepo

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/zura-t/observer.dev/internal/app/models"
	diaryUsecase "github.com/zura-t/observer.dev/internal/app/usecases/diary"
)

func (r *repo) CreateDiaryEntry(ctx context.Context, entry *diaryUsecase.CreateDiaryEntry) (*models.Diary, error) {
	newEntry := newDiaryEntry(entry)
	if newEntry.EntryDate.IsZero() {
		newEntry.EntryDate = time.Now().UTC()
	}
	sql, args, err := squirrel.Insert(tableDiaryEntries).
		Columns(createEntryColumns...).
		Values(newEntry.Values(createEntryColumns...)...).
		Suffix("Returning *").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("squirrel: %w", err)
	}

	var entryCreated models.Diary
	err = r.db.QueryRow(ctx, sql, args...).Scan(
		&entryCreated.ID,
		&entryCreated.Title,
		&entryCreated.Text,
		&entryCreated.EntryDate,
		&entryCreated.UserID,
		&entryCreated.CreatedAt,
		&entryCreated.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("postgres: %w", convertPGError(err))
	}

	return &entryCreated, nil
}
