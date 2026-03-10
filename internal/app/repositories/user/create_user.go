package userRepo

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/zura-t/observer.dev/internal/app/models"
)

func (r *repo) CreateUser(ctx context.Context, user *models.UserDB) (*models.UserDB, error) {
	newUser := newUser(user)
	sql, args, err := squirrel.Insert(tableUsers).Columns(createUserColumns...).Values(newUser.Values(createUserColumns...)...).Suffix("Returning *").PlaceholderFormat(squirrel.Dollar).ToSql()

	if err != nil {
		return nil, fmt.Errorf("squirrel: %w", err)
	}

	var userCreated models.UserDB
	err = r.db.QueryRow(ctx, sql, args...).Scan(&userCreated.ID, &userCreated.Email, &userCreated.Password, &userCreated.Name, &userCreated.IsVerified, &userCreated.CreatedAt, &userCreated.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("postgres: %w", convertPGError(err))
	}

	return &userCreated, nil
}
