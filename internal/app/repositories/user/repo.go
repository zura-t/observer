package userRepo

import (
	"github.com/jackc/pgx/v5"
	userUsecase "github.com/zura-t/observer.dev/internal/app/usecases/user"
)

type repo struct {
	db *pgx.Conn
}

var (
	_ userUsecase.UserRepositoryInterface = (*repo)(nil)
)

func New(connection *pgx.Conn) *repo {
	return &repo{
		db: connection,
	}
}
