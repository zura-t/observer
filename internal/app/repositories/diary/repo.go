package diaryRepo

import (
	"github.com/jackc/pgx/v5"
	diaryUsecase "github.com/zura-t/observer.dev/internal/app/usecases/diary"
	"github.com/zura-t/observer.dev/pkg/logger"
)

type repo struct {
	db     *pgx.Conn
	logger *logger.Logger
}

var (
	_ diaryUsecase.DiaryRepositoryInterface = (*repo)(nil)
)

func New(connection *pgx.Conn, logger *logger.Logger) *repo {
	return &repo{
		db:     connection,
		logger: logger,
	}
}
