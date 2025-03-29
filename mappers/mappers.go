package mappers

import (
	"time"

	"github.com/GuilhermeFujita/nlw_notes_api/database/model"
	"github.com/GuilhermeFujita/nlw_notes_api/dto"
)

func ToNoteModel(note dto.NoteRequestDTO) model.Note {
	now := time.Now()
	return model.Note{
		Content:  note.Content,
		NoteDate: &now,
	}
}
