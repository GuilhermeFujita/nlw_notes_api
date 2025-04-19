package testdata

import (
	"time"

	"github.com/GuilhermeFujita/nlw_notes_api/database/model"
)

func ExpectedNotes() []model.Note {
	return []model.Note{
		{
			ID:       4,
			Content:  "Test single Note",
			NoteDate: time.Date(2025, time.October, 20, 15, 30, 0, 0, time.UTC),
		},
		{
			ID:       2,
			Content:  "bar",
			NoteDate: time.Date(2025, time.April, 18, 11, 0, 0, 0, time.UTC),
		},
		{
			ID:       1,
			Content:  "foo",
			NoteDate: time.Date(2025, time.April, 18, 10, 0, 0, 0, time.UTC),
		},
		{
			ID:       3,
			Content:  "Test content",
			NoteDate: time.Date(2020, time.August, 6, 10, 0, 0, 0, time.UTC),
		},
	}
}

func ExpectedFilteredNotes() []model.Note {
	return []model.Note{
		{
			ID:       3,
			Content:  "Test content",
			NoteDate: time.Date(2020, time.August, 6, 10, 0, 0, 0, time.UTC),
		},
	}
}

func ExpectedSingleNote() model.Note {
	return model.Note{
		ID:       4,
		Content:  "Test single Note",
		NoteDate: time.Date(2025, time.October, 20, 15, 30, 0, 0, time.UTC),
	}
}
