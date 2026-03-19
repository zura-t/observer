package diaryUsecase

import (
	"context"

	"github.com/zura-t/observer.dev/internal/app/models"
	"github.com/zura-t/observer.dev/internal/config"
)

type DiaryUsecase interface {
	GetEntry(c context.Context, id uint64) (*models.Diary, error)
	GetEntries(c context.Context, fil *DiarySearchFilter) (*[]models.Diary, error)
	CreateDiaryEntry(c context.Context, entry *CreateDiaryEntry) (*models.Diary, error)
	UpdateDiaryEntry(c context.Context, entry *UpdateDiaryEntry) (*models.Diary, error)
	DeleteEntry(c context.Context, id uint64) error
}

var (
	_ DiaryUsecase = (*usecase)(nil)
)

type usecase struct {
	diaryRepo DiaryRepositoryInterface
	config    *config.Config
}

func New(diaryRepo DiaryRepositoryInterface, config *config.Config) *usecase {
	return &usecase{
		diaryRepo,
		config,
	}
}
