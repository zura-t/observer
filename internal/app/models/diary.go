package models

import "time"

type Diary struct {
	ID        uint64
	Title     string
	Text      string
	EntryDate time.Time
	UserID    uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}
