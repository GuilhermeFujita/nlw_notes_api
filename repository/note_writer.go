package repository

import (
	"github.com/GuilhermeFujita/nlw_notes_api/database/model"
	"github.com/GuilhermeFujita/nlw_notes_api/dto"
	"github.com/GuilhermeFujita/nlw_notes_api/mappers"
	"gorm.io/gorm"
)

type (
	NoteRepo struct {
		db *gorm.DB
	}
)

func NewNoteWriter(db *gorm.DB) NoteRepo {
	return NoteRepo{db: db}
}

func (r NoteRepo) SaveNote(note dto.NoteRequestDTO) (model.Note, error) {
	noteModel := mappers.ToNoteModel(note)
	result := r.db.Create(&noteModel)
	if result.Error != nil {
		return model.Note{}, result.Error
	}
	return noteModel, nil
}
