package habitsRepo

import (
	"github.com/jackc/pgx/v5"
	habitsUsecase "github.com/zura-t/observer.dev/internal/app/usecases/habits"
	"github.com/zura-t/observer.dev/pkg/logger"
)

type repo struct {
	db     *pgx.Conn
	logger *logger.Logger
}

var (
	_ habitsUsecase.HabitsRepositoryInterface = (*repo)(nil)
)

func New(connection *pgx.Conn, logger *logger.Logger) *repo {
	return &repo{
		db:     connection,
		logger: logger,
	}
}
