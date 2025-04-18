package repository

import (
	"fmt"
	"strings"

	"github.com/GuilhermeFujita/nlw_notes_api/database/model"
	"gorm.io/gorm"
)

type NotesReader struct {
	db *gorm.DB
}

func NewNoteReader(db *gorm.DB) NotesReader {
	return NotesReader{
		db: db,
	}
}

func (r NotesReader) GetNotes(searchedNote string) ([]model.Note, error) {
	var notes []model.Note
	query := r.db.Model(&model.Note{})

	if searchedNote != "" {
		query = query.Where("LOWER(content) LIKE ?", fmt.Sprintf("%%%s%%", strings.ToLower(searchedNote)))
	}

	err := query.
		Order("note_date desc").
		Find(&notes).Error

	return notes, err
}

func (r NotesReader) GetNote(id int) (model.Note, error) {
	var note model.Note
	err := r.db.First(&note, id).Error
	if err != nil {
		return model.Note{}, err
	}
	return note, nil
}
