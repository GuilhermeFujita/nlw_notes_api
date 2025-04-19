package mappers

import (
	"github.com/GuilhermeFujita/nlw_notes_api/database/model"
	"github.com/GuilhermeFujita/nlw_notes_api/dto"
)

func ToNoteModel(note dto.NoteRequestDTO) model.Note {
	return model.Note{
		Content: note.Content,
	}
}

func MapNotesToDTO(notes []model.Note) []dto.NoteResponseDTO {
	notesResponse := make([]dto.NoteResponseDTO, 0, len(notes))

	for _, note := range notes {
		notesResponse = append(notesResponse, dto.NoteResponseDTO{
			ID:       note.ID,
			Content:  note.Content,
			NoteDate: note.NoteDate,
		})
	}

	return notesResponse
}
