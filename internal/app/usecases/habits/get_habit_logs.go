package habitsUsecase

import (
	"context"
	"fmt"

	"github.com/zura-t/observer.dev/internal/app/models"
)

func (u *usecase) GetHabitLogs(ctx context.Context, filter *GetHabitLogsFilter) (*[]models.HabitLog, error) {
	habitLogs, err := u.habitsRepo.GetHabitLogs(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("get habit logs: %w", err)
	}

	return habitLogs, nil
}
