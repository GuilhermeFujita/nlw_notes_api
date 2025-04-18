package entrypoint

import (
	"net/http"

	"github.com/GuilhermeFujita/nlw_notes_api/database/model"
	"github.com/GuilhermeFujita/nlw_notes_api/mappers"
	"github.com/go-chi/render"
)

type (
	NotesFinder interface {
		GetNotes(searchedNote string) ([]model.Note, error)
	}

	GetNotesHandler struct {
		finder NotesFinder
	}
)

func NewGetNotesHandler(f NotesFinder) GetNotesHandler {
	return GetNotesHandler{finder: f}
}

func (h GetNotesHandler) GetNotes(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")

	notes, err := h.finder.GetNotes(search)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	notesDTO := mappers.MapNotesToDTO(notes)

	render.JSON(w, r, notesDTO)

}
