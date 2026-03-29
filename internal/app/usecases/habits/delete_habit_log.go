package habitsUsecase

import (
	"context"
	"fmt"
)

func (u *usecase) DeleteHabitLog(ctx context.Context, id uint64, userID uint64) error {
	if err := u.habitsRepo.DeleteHabitLog(ctx, id, userID); err != nil {
		return fmt.Errorf("delete habit log: %w", err)
	}

	return nil
}
