package main

import (
	"context"
	"log"
	"net/http"

	"github.com/GuilhermeFujita/nlw_notes_api/cmd/modules"
	"github.com/go-chi/cors"
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
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})

	// 2) envolva o seu router com o middleware de CORS
	handler := corsMiddleware.Handler(router)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Println("Server is starting at :9090")
				http.ListenAndServe(":9090", handler)
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Server stopping...")
			return nil
		},
	})
}
