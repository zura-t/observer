package diaryUsecase

import (
	"context"

	"github.com/zura-t/observer.dev/internal/app/models"
)

type DiaryRepositoryInterface interface {
	GetEntry(c context.Context, id uint64, userID uint64) (*models.Diary, error)
	GetEntries(c context.Context, fil *DiarySearchFilter) (*[]models.Diary, error)
	CreateDiaryEntry(c context.Context, entry *CreateDiaryEntry) (*models.Diary, error)
	UpdateDiaryEntry(c context.Context, entry *UpdateDiaryEntry) (*models.Diary, error)
	DeleteEntry(c context.Context, id uint64, userID uint64) error
}
