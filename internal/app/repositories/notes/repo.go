package notesRepo

import (
	"github.com/jackc/pgx/v5"
	notesUsecase "github.com/zura-t/observer.dev/internal/app/usecases/notes"
	"github.com/zura-t/observer.dev/pkg/logger"
)

type repo struct {
	db     *pgx.Conn
	logger *logger.Logger
}

var (
	_ notesUsecase.NotesRepositoryInterface = (*repo)(nil)
)

func New(connection *pgx.Conn, logger *logger.Logger) *repo {
	return &repo{
		db:     connection,
		logger: logger,
	}
}
