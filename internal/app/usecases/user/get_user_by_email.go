package userUsecase

import (
	"context"
	"fmt"

	"github.com/zura-t/observer.dev/internal/app/models"
)

func (u *usecase) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	userDB, err := u.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("get user by email: %w", err)
	}

	resp := UserDBToUser(userDB)

	return resp, nil
}
