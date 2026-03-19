package diaryUsecase

import (
	"context"
	"fmt"
)

func (u *usecase) DeleteEntry(ctx context.Context, id uint64) error {
	if err := u.diaryRepo.DeleteEntry(ctx, id); err != nil {
		return fmt.Errorf("delete entry: %w", err)
	}

	return nil
}
