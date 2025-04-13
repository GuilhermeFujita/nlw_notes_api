package entrypoint

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type (
	NoteDeleter interface {
		DeleteNote(noteID int) error
	}

	DeleteNoteHandler struct {
		deleter NoteDeleter
	}
)

func NewDeleteNoteHandler(deleter NoteDeleter) DeleteNoteHandler {
	return DeleteNoteHandler{
		deleter: deleter,
	}
}

func (h DeleteNoteHandler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	noteIDString := chi.URLParam(r, "id")

	noteID, err := strconv.Atoi(noteIDString)
	if err != nil {
		http.Error(w, "Invalid note id", http.StatusBadRequest)
		return
	}

	if err := h.deleter.DeleteNote(noteID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Note not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
