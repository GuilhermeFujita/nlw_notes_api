package usecase

import (
	"github.com/GuilhermeFujita/nlw_notes_api/database/model"
	"github.com/GuilhermeFujita/nlw_notes_api/dto"
)

type (
	NoteSaver interface {
		SaveNote(note dto.NoteRequestDTO) (model.Note, error)
	}

	NoteUseCase struct {
		saver NoteSaver
	}
)

func NewNoteUseCase(s NoteSaver) NoteUseCase {
	return NoteUseCase{saver: s}
}

func (u NoteUseCase) CreateNote(note dto.NoteRequestDTO) (model.Note, error) {
	return u.saver.SaveNote(note)
}
