package userUsecase

import (
	"time"

	"github.com/zura-t/observer.dev/internal/app/models"
)

func UserDBToUser(userDB *models.UserDB) *models.User {
	return &models.User{
		ID:         userDB.ID,
		Name:       userDB.Name,
		Email:      userDB.Email,
		IsVerified: userDB.IsVerified,
		CreatedAt:  userDB.CreatedAt,
		UpdatedAt:  userDB.UpdatedAt,
	}
}

func UserDBToUserWithToken(userDB *models.UserDB, accessToken string, accessTokenExpiresAt time.Time, refreshToken string, refreshTokenExpiresAt time.Time) *models.UserWithToken {
	return &models.UserWithToken{
		User:                  UserDBToUser(userDB),
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessTokenExpiresAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshTokenExpiresAt,
	}
}
