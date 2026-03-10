package userUsecase

import (
	"context"
	"fmt"

	"github.com/zura-t/observer.dev/internal/app/models"
)

func (u *usecase) GetUserByID(ctx context.Context, id uint64) (*models.User, error) {
	userDB, err := u.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get user by id: %w", err)
	}

	resp := UserDBToUser(userDB)

	return resp, nil
}
