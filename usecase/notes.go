package usecase

import (
	"github.com/GuilhermeFujita/nlw_notes_api/database/model"
	"github.com/GuilhermeFujita/nlw_notes_api/dto"
)

type (
	NoteSaver interface {
		SaveNote(note dto.NoteRequestDTO) (model.Note, error)
	}

	NotesFinder interface {
		GetNotes(searchedNote string) ([]model.Note, error)
	}

	NoteUseCase struct {
		saver  NoteSaver
		finder NotesFinder
	}
)

func NewNoteUseCase(s NoteSaver, f NotesFinder) NoteUseCase {
	return NoteUseCase{
		saver:  s,
		finder: f,
	}
}

func (u NoteUseCase) CreateNote(note dto.NoteRequestDTO) (model.Note, error) {
	return u.saver.SaveNote(note)
}

func (u NoteUseCase) GetNotes(searchedNote string) ([]model.Note, error) {
	return u.finder.GetNotes(searchedNote)
}
