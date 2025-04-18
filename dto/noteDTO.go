package dto

import "time"

type NoteRequestDTO struct {
	Content string `json:"content"`
}

type NoteResponseDTO struct {
	ID       int64     `json:"id"`
	NoteDate time.Time `json:"noteDate"`
	Content  string    `json:"content"`
}
