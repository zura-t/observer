package diaryUsecase

import (
	"context"
	"fmt"

	"github.com/zura-t/observer.dev/internal/app/models"
)

func (u *usecase) CreateDiaryEntry(ctx context.Context, entry *CreateDiaryEntry) (*models.Diary, error) {
	diaryCreated, err := u.diaryRepo.CreateDiaryEntry(ctx, entry)
	if err != nil {
		return nil, fmt.Errorf("create entry: %w", err)
	}

	return diaryCreated, nil
}
