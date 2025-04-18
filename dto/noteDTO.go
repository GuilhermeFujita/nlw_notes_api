package dto

import (
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type NoteRequestDTO struct {
	Content string `json:"content" validate:"required,min=3,max=255"`
}

type NoteResponseDTO struct {
	ID       int64     `json:"id"`
	NoteDate time.Time `json:"noteDate"`
	Content  string    `json:"content"`
}

func (p NoteRequestDTO) Validate() error {
	p.Content = strings.TrimSpace(p.Content)
	validate := validator.New()
	return validate.Struct(p)
}
