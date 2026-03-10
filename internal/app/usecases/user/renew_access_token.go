package userUsecase

import (
	"context"

	"github.com/zura-t/observer.dev/internal/app/models"
)

func (u *usecase) RenewAccessToken(ctx context.Context, refreshToken string) (*NewAccessToken, error) {
	refreshPayload, err := u.tokenMaker.VerifyToken(refreshToken)
	if err != nil {
		return nil, err
	}

	accessTokenPayload := NewPayload(&models.UserDB{
		ID:         refreshPayload.ID,
		Email:      refreshPayload.Email,
		Name:       refreshPayload.Name,
		IsVerified: refreshPayload.IsVerified,
	}, u.config.AccessTokenDuration)

	accessToken, _, err := u.tokenMaker.CreateToken(accessTokenPayload)
	if err != nil {
		return nil, err
	}

	resp := &NewAccessToken{
		AccessToken: accessToken,
	}

	return resp, nil
}
