package model

import "time"

type Note struct {
	ID        int64 `gorm:"primaryKey"`
	NoteDate  *time.Time
	Content   string `gorm:"not null;size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
