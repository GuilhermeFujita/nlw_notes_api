package entrypoint

import (
	"encoding/json"
	"net/http"

	"github.com/GuilhermeFujita/nlw_notes_api/database/model"
	"github.com/GuilhermeFujita/nlw_notes_api/dto"
)

type (
	NoteCreator interface {
		CreateNote(note dto.NoteRequestDTO) (model.Note, error)
	}

	NoteHandler struct {
		creator NoteCreator
	}
)

func NewNoteHandler(c NoteCreator) NoteHandler {
	return NoteHandler{c}
}

func (h NoteHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var p dto.NoteRequestDTO

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	note, err := h.creator.CreateNote(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(note); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
