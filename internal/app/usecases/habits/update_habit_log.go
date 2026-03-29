package habitsUsecase

import (
	"context"
	"fmt"

	"github.com/zura-t/observer.dev/internal/app/models"
)

func (u *usecase) UpdateHabitLog(ctx context.Context, habitLog *UpdateHabitLog) (*models.HabitLog, error) {
	habitLogUpdated, err := u.habitsRepo.UpdateHabitLog(ctx, habitLog)
	if err != nil {
		return nil, fmt.Errorf("update habit log: %w", err)
	}

	return habitLogUpdated, nil
}
