package router

import (
	"net/http"

	"github.com/GuilhermeFujita/nlw_notes_api/entrypoint"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(
	createHandler entrypoint.CreateNoteHandler,
	getHandler entrypoint.GetNotesHandler,
	deleteHandler entrypoint.DeleteNoteHandler,
) http.Handler {

	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	router.Route("/notes", func(r chi.Router) {
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Pong"))
		})
		r.Get("/", getHandler.GetNotes)
		r.Post("/", createHandler.CreateNote)
		r.Delete("/{id}", deleteHandler.DeleteNote)
	})

	return router
}
