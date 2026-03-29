package habitsUsecase

import (
	"context"
	"fmt"

	"github.com/zura-t/observer.dev/internal/app/models"
)

func (u *usecase) GetHabitLog(ctx context.Context, id uint64, userID uint64) (*models.HabitLog, error) {
	habitLog, err := u.habitsRepo.GetHabitLog(ctx, id, userID)
	if err != nil {
		return nil, fmt.Errorf("get habit log: %w", err)
	}

	return habitLog, nil
}
