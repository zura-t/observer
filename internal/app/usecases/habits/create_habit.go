package habitsUsecase

import (
	"context"
	"fmt"

	"github.com/zura-t/observer.dev/internal/app/models"
)

func (u *usecase) CreateHabit(ctx context.Context, habit *CreateHabit) (*models.Habit, error) {
	habitCreated, err := u.habitsRepo.CreateHabit(ctx, habit)
	if err != nil {
		return nil, fmt.Errorf("create habit: %w", err)
	}

	return habitCreated, nil
}
