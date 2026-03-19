package diaryController

type CreateDiaryRequest struct {
	Title     string `json:"title" binding:"required"`
	Text      string `json:"text" binding:"required"`
	EntryDate string `json:"entry_date"`
}

type GetDiaryByIDRequest struct {
	ID uint64 `uri:"id" binding:"required"`
}

type GetDiaryRequest struct {
	DateFrom string `form:"date_from"`
	DateTo   string `form:"date_to"`
}

type UpdateDiaryRequest struct {
	Title     string `json:"title"`
	Text      string `json:"text"`
	EntryDate string `json:"entry_date"`
}

type DeleteDiaryRequest struct {
	ID uint64 `uri:"id" binding:"required"`
}
