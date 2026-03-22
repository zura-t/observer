package diaryController

import "time"

type CreateDiaryRequest struct {
	Title     string    `json:"title" binding:"required"`
	Text      string    `json:"text" binding:"required"`
	EntryDate time.Time `json:"entry_date"`
}

type GetDiaryByIDRequest struct {
	ID uint64 `uri:"id" binding:"required"`
}

type GetDiaryRequest struct {
	Limit  uint32 `form:"limit,default=10"`
	Offset uint32 `form:"offset,default=0"`
}

type UpdateDiaryRequest struct {
	Title     string `json:"title"`
	Text      string `json:"text"`
	EntryDate string `json:"entry_date"`
}

type DeleteDiaryRequest struct {
	ID uint64 `uri:"id" binding:"required"`
}
