package testdata

import (
	"time"

	"github.com/GuilhermeFujita/nlw_notes_api/database/model"
	"github.com/GuilhermeFujita/nlw_notes_api/dto"
)

func NoteToCreate() dto.NoteRequestDTO {
	return dto.NoteRequestDTO{
		Content: "Test note",
	}
}

func ExpectedNoteCreated() model.Note {
	return model.Note{
		ID:      1,
		Content: "Test note",
	}
}

func NotesWithoutFiltersResult() []model.Note {
	return []model.Note{
		{
			ID:       1,
			Content:  "Note with search 1",
			NoteDate: time.Date(2025, time.April, 20, 10, 30, 0, 0, time.UTC),
		},
		{
			ID:       2,
			Content:  "This note has search on content",
			NoteDate: time.Date(2024, time.October, 10, 15, 45, 0, 0, time.UTC),
		},
		{
			ID:       3,
			Content:  "Learn Golang test",
			NoteDate: time.Date(2025, time.April, 20, 15, 50, 0, 0, time.UTC),
		},
	}
}

func FilteredNotesResult() []model.Note {
	return []model.Note{
		{
			ID:       1,
			Content:  "Note with search 1",
			NoteDate: time.Date(2025, time.April, 20, 10, 30, 0, 0, time.UTC),
		},
		{
			ID:       2,
			Content:  "This note has search on content",
			NoteDate: time.Date(2024, time.October, 10, 15, 45, 0, 0, time.UTC),
		},
	}
}

func ExpectedNoteToDelete() model.Note {
	return model.Note{
		ID:       1,
		Content:  "Note to delete",
		NoteDate: time.Date(2025, time.December, 12, 10, 0, 0, 0, time.UTC),
	}
}
