package diaryUsecase

import (
	"context"
	"fmt"

	"github.com/zura-t/observer.dev/internal/app/models"
)

func (u *usecase) GetEntries(ctx context.Context, filter *DiarySearchFilter) (*[]models.Diary, error) {
	entries, err := u.diaryRepo.GetEntries(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("get entries: %w", err)
	}

	return entries, nil
}
