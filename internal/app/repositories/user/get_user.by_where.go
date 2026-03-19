package userRepo

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/zura-t/observer.dev/internal/app/models"
	userUsecase "github.com/zura-t/observer.dev/internal/app/usecases/user"
)

func (r *repo) GetUserByWhere(ctx context.Context, filter *userUsecase.UserSearchFilter) (*models.UserDB, error) {
	qb := squirrel.Select(tableUsersColumns...).From(tableUsers).PlaceholderFormat(squirrel.Dollar)
	qb = applyUserSearchFilter(qb, filter)
	sql, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("squirrel: %w", err)
	}

	var users models.UserDB
	err = r.db.QueryRow(ctx, sql, args...).Scan(&users)
	if err != nil {
		return nil, fmt.Errorf("postgres: %w", convertPGError(err))
	}

	return &users, nil
}

func applyUserSearchFilter(qb squirrel.SelectBuilder, filter *userUsecase.UserSearchFilter) squirrel.SelectBuilder {
	if filter == nil {
		return qb
	}

	if filter.ID != nil {
		qb = qb.Where(squirrel.Eq{columnID: filter.ID})
	}

	if filter.Email != "" {
		qb = qb.Where(squirrel.Eq{columnEmail: filter.Email})
	}

	if filter.Name != "" {
		qb = qb.Where(squirrel.Eq{columnName: filter.Name})
	}

	if filter.IsVerified != nil {
		qb = qb.Where(squirrel.Eq{columnIsVerified: filter.IsVerified})
	}

	return qb
}
