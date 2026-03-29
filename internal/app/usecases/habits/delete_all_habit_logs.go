package habitsUsecase

import (
	"context"
	"fmt"
)

func (u *usecase) DeleteAllHabitLogs(ctx context.Context, userID uint64) error {
	if err := u.habitsRepo.DeleteAllHabitLogs(ctx, userID); err != nil {
		return fmt.Errorf("delete all habit logs: %w", err)
	}

	return nil
}
