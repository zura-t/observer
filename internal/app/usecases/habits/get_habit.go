package habitsUsecase

import (
	"context"
	"fmt"

	"github.com/zura-t/observer.dev/internal/app/models"
)

func (u *usecase) GetHabit(ctx context.Context, id uint64, userID uint64) (*models.Habit, error) {
	habit, err := u.habitsRepo.GetHabit(ctx, id, userID)
	if err != nil {
		return nil, fmt.Errorf("get habit: %w", err)
	}

	return habit, nil
}
