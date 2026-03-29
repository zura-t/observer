package habitsUsecase

import (
	"context"
	"fmt"

	"github.com/zura-t/observer.dev/internal/app/models"
)

func (u *usecase) UpdateHabit(ctx context.Context, habit *UpdateHabit) (*models.Habit, error) {
	habitUpdated, err := u.habitsRepo.UpdateHabit(ctx, habit)
	if err != nil {
		return nil, fmt.Errorf("update habit: %w", err)
	}

	return habitUpdated, nil
}
