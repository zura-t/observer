package habitsUsecase

import (
	"context"
	"fmt"
)

func (u *usecase) DeleteHabitLogs(ctx context.Context, habitID uint64, userID uint64) error {
	if err := u.habitsRepo.DeleteHabitLogs(ctx, habitID, userID); err != nil {
		return fmt.Errorf("delete habit logs: %w", err)
	}

	return nil
}
