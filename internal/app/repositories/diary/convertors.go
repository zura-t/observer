package diaryRepo

import (
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/zura-t/observer.dev/internal/app/models"
	diaryUsecase "github.com/zura-t/observer.dev/internal/app/usecases/diary"
)

func newDiaryEntry(o *diaryUsecase.CreateDiaryEntry) *diaryEntry {
	return &diaryEntry{
		Title:     o.Title,
		Text:      o.Text,
		EntryDate: o.EntryDate,
		UserID:    o.UserID,
	}
}

func convertPGError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return models.ErrDiaryEntryNotFound
	}
	return err
}
