package userRepo

import (
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/zura-t/observer.dev/internal/app/models"
)

func newUser(o *models.UserDB) *user {
	return &user{
		Email:    o.Email,
		Password: o.Password,
		Name:     o.Name,
	}
}

func convertPGError(err error) error {
	if err == nil {
		return nil
	}

	if errors.As(err, &pgx.ErrNoRows) {
		return models.ErrUserNotFound
	}
	return err
}
