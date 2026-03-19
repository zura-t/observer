package userRepo

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/zura-t/observer.dev/internal/app/models"
	userUsecase "github.com/zura-t/observer.dev/internal/app/usecases/user"
)

func (r *repo) UpdateUser(ctx context.Context, id uint64, user *userUsecase.UpdateUser) (*models.UserDB, error) {
	qb := squirrel.Update(tableUsers).Where(squirrel.Eq{columnID: id}).Suffix("Returning *").PlaceholderFormat(squirrel.Dollar)
	qb = applyUserUpdate(qb, user)
	sql, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("squirrel: %w", err)
	}

	var userUpdated models.UserDB
	err = r.db.QueryRow(ctx, sql, args...).Scan(&userUpdated.ID, &userUpdated.Email, &userUpdated.Password, &userUpdated.Name, &userUpdated.IsVerified, &userUpdated.CreatedAt, &userUpdated.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("postgres: %w", convertPGError(err))
	}

	return &userUpdated, nil
}

func applyUserUpdate(qb squirrel.UpdateBuilder, user *userUsecase.UpdateUser) squirrel.UpdateBuilder {
	if user.Email != "" {
		qb = qb.Set(columnEmail, user.Email)
	}

	if user.Name != "" {
		qb = qb.Set(columnName, user.Name)
	}

	if user.IsVerified != nil {
		qb = qb.Set(columnIsVerified, user.IsVerified)
	}

	qb = qb.Set(columnUpdatedAt, squirrel.Expr("NOW()"))

	return qb
}
