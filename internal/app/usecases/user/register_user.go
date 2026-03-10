package userUsecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/zura-t/observer.dev/internal/app/models"
	"github.com/zura-t/observer.dev/pkg/token"
	"github.com/zura-t/observer.dev/pkg/utils"
)

func (u *usecase) RegisterUser(ctx context.Context, user *RegisterUser) (*models.UserWithToken, error) {
	// ? get user account
	userExists, err := u.userRepo.GetUserByEmail(ctx, user.Email)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("register user: %w", err)
	}

	if userExists != nil {
		return nil, models.ErrEmailAlreadyExists
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, fmt.Errorf("register user: %w", err)
	}

	// ? send email

	userCreated, err := u.userRepo.CreateUser(ctx, &models.UserDB{
		Email:    user.Email,
		Password: hashedPassword,
		Name:     user.Name,
	})
	if err != nil {
		return nil, fmt.Errorf("register user: %w", err)
	}

	accessTokenPayload := NewPayload(userCreated, u.config.AccessTokenDuration)

	accessToken, _, err := u.tokenMaker.CreateToken(accessTokenPayload)
	if err != nil {
		return nil, fmt.Errorf("register user: %w", err)
	}

	refreshTokenPayload := NewPayload(userCreated, u.config.RefreshTokenDuration)
	
	refreshToken, _, err := u.tokenMaker.CreateToken(refreshTokenPayload)
	if err != nil {
		return nil, fmt.Errorf("register user: %w", err)
	}

	resp := UserDBToUserWithToken(userCreated, accessToken, accessTokenPayload.ExpiredAt, refreshToken, refreshTokenPayload.ExpiredAt)

	return resp, nil
}

func NewPayload(user *models.UserDB, duration time.Duration) *token.Payload {
	return &token.Payload{
		ID:         user.ID,
		Email:      user.Email,
		Name:       user.Name,
		IsVerified: user.IsVerified,
		IssuedAt:   time.Now(),
		ExpiredAt:  time.Now().Add(duration),
	}
}
