package userUsecase

import (
	"context"
	"fmt"

	"github.com/zura-t/observer.dev/internal/app/models"
)

func (u *usecase) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	userDB, err := u.userRepo.GetUserByWhere(ctx, &models.UserSearchFilter{
		Name: username,
	})
	if err != nil {
		return nil, fmt.Errorf("get user by username: %w", err)
	}

	resp := UserDBToUser(userDB)

	return resp, nil
}
