package main

import (
	"context"
	"log"
	"net/http"

	"github.com/GuilhermeFujita/nlw_notes_api/cmd/modules"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(

		modules.AppModule,
		fx.Invoke(StartServer),
	)
	app.Run()
}

func StartServer(lc fx.Lifecycle, router http.Handler) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Println("Server is starting at :9090")
				http.ListenAndServe(":9090", router)
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Server stopping...")
			return nil
		},
	})
}
