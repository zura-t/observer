package diaryRepo

import (
	"github.com/jackc/pgx/v5"
	diaryUsecase "github.com/zura-t/observer.dev/internal/app/usecases/diary"
)

type repo struct {
	db *pgx.Conn
}

var (
	_ diaryUsecase.DiaryRepositoryInterface = (*repo)(nil)
)

func New(connection *pgx.Conn) *repo {
	return &repo{
		db: connection,
	}
}