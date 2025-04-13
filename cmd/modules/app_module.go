package modules

import (
	"github.com/GuilhermeFujita/nlw_notes_api/database"
	"github.com/GuilhermeFujita/nlw_notes_api/entrypoint"
	"github.com/GuilhermeFujita/nlw_notes_api/repository"
	"github.com/GuilhermeFujita/nlw_notes_api/router"
	"github.com/GuilhermeFujita/nlw_notes_api/usecase"
	"go.uber.org/fx"
)

var AppModule = fx.Options(
	// Constructors
	fx.Provide(
		database.Connect,
		repository.NewNoteWriter,
		repository.NewNoteReader,
		usecase.NewNoteUseCase,
		entrypoint.NewCreateNoteHandler,
		entrypoint.NewGetNotesHandler,
		entrypoint.NewDeleteNoteHandler,
		router.NewRouter,
	),

	//UseCases
	fx.Provide(
		func(r repository.NotesReader) usecase.NotesFinder { return r },
		func(w repository.NotesWriter) usecase.NoteWriter { return w },
	),

	//Entrypoint
	fx.Provide(
		func(u usecase.NoteUseCase) entrypoint.NotesFinder { return u },
		func(u usecase.NoteUseCase) entrypoint.NoteCreator { return u },
		func(u usecase.NoteUseCase) entrypoint.NoteDeleter { return u },
	),
)
