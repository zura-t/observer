package habitsUsecase

import (
	"context"
	"fmt"
)

func (u *usecase) DeleteHabit(ctx context.Context, id uint64, userID uint64) error {
	if err := u.habitsRepo.DeleteHabit(ctx, id, userID); err != nil {
		return fmt.Errorf("delete habit: %w", err)
	}

	return nil
}
