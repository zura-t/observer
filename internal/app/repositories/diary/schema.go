package diaryRepo

import (
	"strings"
	"time"
)

const tableDiaryEntries = "public.diary_entries"

const (
	columnID        = "id"
	columnTitle     = "title"
	columnText      = "text"
	columnEntryDate = "entry_date"
	columnUserID    = "user_id"
	columnCreatedAt = "created_at"
	columnUpdatedAt = "updated_at"
)

type diaryEntry struct {
	Title     string    `db:"title"`
	Text      string    `db:"text"`
	EntryDate time.Time `db:"entry_date"`
	UserID    uint64    `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

var (
	tableDiaryEntriesColumns = []string{
		columnID,
		columnTitle,
		columnText,
		columnEntryDate,
		columnUserID,
		columnCreatedAt,
		columnUpdatedAt,
	}

	createEntryColumns = []string{
		columnTitle,
		columnText,
		columnUserID,
	}
)

func (o *diaryEntry) mapFields() map[string]any {
	return map[string]any{
		columnTitle:     o.Title,
		columnText:      o.Text,
		columnEntryDate: o.EntryDate,
		columnUserID:    o.UserID,
		columnCreatedAt: o.CreatedAt,
		columnUpdatedAt: o.UpdatedAt,
	}
}

func (o *diaryEntry) Values(columns ...string) []any {
	mapFields := o.mapFields()
	values := make([]any, 0, len(columns))
	for i := range columns {
		if v, ok := mapFields[columns[i]]; ok {
			values = append(values, v)
		} else {
			values = append(values, nil)
		}
	}
	return values
}

func (o *diaryEntry) ReturningValues(columns ...string) string {
	return strings.Join(columns, ", ")
}
