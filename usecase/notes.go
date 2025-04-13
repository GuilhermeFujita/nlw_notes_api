package usecase

import (
	"github.com/GuilhermeFujita/nlw_notes_api/database/model"
	"github.com/GuilhermeFujita/nlw_notes_api/dto"
)

type (
	NoteWriter interface {
		SaveNote(note dto.NoteRequestDTO) (model.Note, error)
		DeleteNote(note model.Note) error
	}

	NotesFinder interface {
		GetNotes(searchedNote string) ([]model.Note, error)
		GetNote(id int) (model.Note, error)
	}

	NoteUseCase struct {
		writer NoteWriter
		finder NotesFinder
	}
)

func NewNoteUseCase(w NoteWriter, f NotesFinder) NoteUseCase {
	return NoteUseCase{
		writer: w,
		finder: f,
	}
}

func (u NoteUseCase) CreateNote(note dto.NoteRequestDTO) (model.Note, error) {
	return u.writer.SaveNote(note)
}

func (u NoteUseCase) GetNotes(searchedNote string) ([]model.Note, error) {
	return u.finder.GetNotes(searchedNote)
}

func (u NoteUseCase) DeleteNote(noteID int) error {
	noteToDelete, err := u.finder.GetNote(noteID)
	if err != nil {
		return err
	}

	return u.writer.DeleteNote(noteToDelete)
}
