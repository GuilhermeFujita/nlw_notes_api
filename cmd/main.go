package main

import (
	"net/http"

	"github.com/GuilhermeFujita/nlw_notes_api/database"
	"github.com/GuilhermeFujita/nlw_notes_api/entrypoint"
	"github.com/GuilhermeFujita/nlw_notes_api/repository"
	"github.com/GuilhermeFujita/nlw_notes_api/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	noteWriter := repository.NewNoteWriter(db)
	noteUsecase := usecase.NewNoteUseCase(noteWriter)
	noteHandler := entrypoint.NewNoteHandler(noteUsecase)

	router.Route("/notes", func(r chi.Router) {
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Pong"))
		})
		r.Post("/", noteHandler.CreateNote)
	})

	http.ListenAndServe(":9090", router)
}
