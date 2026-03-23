package diaryRepo

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/zura-t/observer.dev/internal/app/models"
)

func (r *repo) GetEntry(ctx context.Context, id uint64, userID uint64) (*models.Diary, error) {
	sql, args, err := squirrel.Select(tableDiaryEntriesColumns...).
		From(tableDiaryEntries).
		Where(squirrel.Eq{columnID: id, columnUserID: userID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("squirrel: %w", err)
	}

	var entry models.Diary
	err = r.db.QueryRow(ctx, sql, args...).Scan(
		&entry.ID,
		&entry.Title,
		&entry.Text,
		&entry.EntryDate,
		&entry.UserID,
		&entry.CreatedAt,
		&entry.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("postgres: %w", convertPGError(err))
	}

	return &entry, nil
}
