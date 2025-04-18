package entrypoint

import (
	"encoding/json"
	"net/http"

	"github.com/GuilhermeFujita/nlw_notes_api/database/model"
	"github.com/GuilhermeFujita/nlw_notes_api/dto"
	"github.com/GuilhermeFujita/nlw_notes_api/mappers"
	"github.com/go-chi/render"
)

type (
	NoteCreator interface {
		CreateNote(note dto.NoteRequestDTO) (model.Note, error)
	}

	CreateNoteHandler struct {
		creator NoteCreator
	}
)

func NewCreateNoteHandler(c NoteCreator) CreateNoteHandler {
	return CreateNoteHandler{c}
}

func (h CreateNoteHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var p dto.NoteRequestDTO

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := p.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	note, err := h.creator.CreateNote(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	notesResponse := mappers.MapNotesToDTO([]model.Note{note})

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, notesResponse)

}
