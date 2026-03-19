package diaryUsecase

import "time"

type CreateDiaryEntry struct {
	ID        uint64
	Title     string
	Text      string
	EntryDate time.Time
	UserID    uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DiarySearchFilter struct {
	UserID   uint64
	DateFrom *time.Time
	DateTo   *time.Time
}

type UpdateDiaryEntry struct {
	ID        uint64
	Title     string
	Text      string
	EntryDate *time.Time
}
