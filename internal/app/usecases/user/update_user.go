package userUsecase

import (
	"context"
	"fmt"

	"github.com/zura-t/observer.dev/internal/app/models"
)

func (u *usecase) UpdateUser(ctx context.Context, id uint64, user *UpdateUser) (*models.User, error) {
	userDB, err := u.userRepo.UpdateUser(ctx, id, user)
	if err != nil {
		return nil, fmt.Errorf("update user: %w", err)
	}

	resp := UserDBToUser(userDB)

	return resp, nil
}
