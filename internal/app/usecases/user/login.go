package userUsecase

import (
	"context"

	"github.com/zura-t/observer.dev/internal/app/models"
	"github.com/zura-t/observer.dev/pkg/utils"
)

func (u *usecase) Login(ctx context.Context, user *Login) (*models.UserWithToken, error) {
	userDB, err := u.userRepo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	err = utils.CheckPassword(user.Password, userDB.Password)
	if err != nil {
		return nil, err
	}

	accessTokenPayload := NewPayload(userDB, u.config.AccessTokenDuration)

	accessToken, _, err := u.tokenMaker.CreateToken(accessTokenPayload)
	if err != nil {
		return nil, err
	}

	refreshTokenPayload := NewPayload(userDB, u.config.RefreshTokenDuration)

	refreshToken, _, err := u.tokenMaker.CreateToken(refreshTokenPayload)
	if err != nil {
		return nil, err
	}

	resp := UserDBToUserWithToken(userDB, accessToken, accessTokenPayload.ExpiredAt, refreshToken, refreshTokenPayload.ExpiredAt)

	return resp, nil
}
