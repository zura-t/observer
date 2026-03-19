package diaryUsecase

import (
	"context"
	"fmt"

	"github.com/zura-t/observer.dev/internal/app/models"
)

func (u *usecase) UpdateDiaryEntry(ctx context.Context, entry *UpdateDiaryEntry) (*models.Diary, error) {
	diaryUpdated, err := u.diaryRepo.UpdateDiaryEntry(ctx, entry)
	if err != nil {
		return nil, fmt.Errorf("update entry: %w", err)
	}

	return diaryUpdated, nil
}
