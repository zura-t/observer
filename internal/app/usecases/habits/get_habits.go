package habitsUsecase

import (
	"context"
	"fmt"

	"github.com/zura-t/observer.dev/internal/app/models"
)

func (u *usecase) GetHabits(ctx context.Context, userID uint64) (*[]models.Habit, error) {
	habits, err := u.habitsRepo.GetHabits(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get habits: %w", err)
	}

	return habits, nil
}
