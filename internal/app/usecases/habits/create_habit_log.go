package habitsUsecase

import (
	"context"
	"fmt"

	"github.com/zura-t/observer.dev/internal/app/models"
)

func (u *usecase) CreateHabitLog(ctx context.Context, habitLog *CreateHabitLog) (*models.HabitLog, error) {
	habitLogCreated, err := u.habitsRepo.CreateHabitLog(ctx, habitLog)
	if err != nil {
		return nil, fmt.Errorf("create habit log: %w", err)
	}

	return habitLogCreated, nil
}
