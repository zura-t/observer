package diaryUsecase

import (
	"context"
	"fmt"

	"github.com/zura-t/observer.dev/internal/app/models"
)

func (u *usecase) GetEntry(ctx context.Context, id uint64) (*models.Diary, error) {
	entry, err := u.diaryRepo.GetEntry(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get entry: %w", err)
	}

	return entry, nil
}
