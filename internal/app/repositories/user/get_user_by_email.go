package userRepo

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/zura-t/observer.dev/internal/app/models"
)

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*models.UserDB, error) {
	sql, args, err := squirrel.Select(tableUsersColumns...).From(tableUsers).Where(squirrel.Eq{columnEmail: email}).PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, fmt.Errorf("squirrel: %w", err)
	}

	user := models.UserDB{}
	err = r.db.QueryRow(ctx, sql, args...).Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.IsVerified, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("postgres: %w", convertPGError(err))
	}

	return &user, nil
}
