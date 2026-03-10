package userUsecase

import (
	"context"

	"github.com/zura-t/observer.dev/internal/app/models"
	"github.com/zura-t/observer.dev/internal/config"
	"github.com/zura-t/observer.dev/pkg/token"
)

type UserUsecase interface {
	RegisterUser(ctx context.Context, user *RegisterUser) (*models.UserWithToken, error)
	Login(ctx context.Context, user *Login) (*models.UserWithToken, error)
	GetUserByID(ctx context.Context, id uint64) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	UpdateUser(ctx context.Context, id uint64, user *UpdateUser) (*models.User, error)
	DeleteUser(ctx context.Context, id uint64) error
	RenewAccessToken(ctx context.Context, refreshToken string) (*NewAccessToken, error)
}

var (
	_ UserUsecase = (*usecase)(nil)
)

type usecase struct {
	userRepo   UserRepositoryInterface
	tokenMaker token.Maker
	config     *config.Config
}

func New(userRepo UserRepositoryInterface, tokenMaker token.Maker, config *config.Config) *usecase {
	return &usecase{
		userRepo,
		tokenMaker,
		config,
	}
}
