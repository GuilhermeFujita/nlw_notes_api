package repository

import (
	"github.com/GuilhermeFujita/nlw_notes_api/database/model"
	"github.com/GuilhermeFujita/nlw_notes_api/dto"
	"github.com/GuilhermeFujita/nlw_notes_api/mappers"
	"gorm.io/gorm"
)

type (
	NotesWriter struct {
		db *gorm.DB
	}
)

func NewNoteWriter(db *gorm.DB) NotesWriter {
	return NotesWriter{db: db}
}

func (w NotesWriter) SaveNote(note dto.NoteRequestDTO) (model.Note, error) {
	noteModel := mappers.ToNoteModel(note)
	result := w.db.Create(&noteModel)
	if result.Error != nil {
		return model.Note{}, result.Error
	}
	return noteModel, nil
}

func (w NotesWriter) DeleteNote(note model.Note) error {
	return w.db.Delete(&note).Error
}
