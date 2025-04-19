package model

import (
	"time"
)

type Note struct {
	ID        int64     `gorm:"primaryKey"`
	NoteDate  time.Time `gorm:"autoCreateTime"`
	Content   string    `gorm:"not null;size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
