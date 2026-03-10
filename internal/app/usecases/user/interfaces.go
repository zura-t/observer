package userUsecase

import (
	"context"

	"github.com/zura-t/observer.dev/internal/app/models"
)

type UserRepositoryInterface interface {
	GetUserByID(ctx context.Context, id uint64) (*models.UserDB, error)
	CreateUser(ctx context.Context, user *models.UserDB) (*models.UserDB, error)
	GetUserByEmail(ctx context.Context, email string) (*models.UserDB, error)
	GetUserByWhere(ctx context.Context, filter *models.UserSearchFilter) (*models.UserDB, error)
	UpdateUser(ctx context.Context, id uint64, user *UpdateUser) (*models.UserDB, error)
	DeleteUser(ctx context.Context, id uint64) error
}
