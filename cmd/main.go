package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/GuilhermeFujita/nlw_notes_api/cmd/modules"
	"github.com/GuilhermeFujita/nlw_notes_api/config"
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

func StartServer(lc fx.Lifecycle, router http.Handler, config *config.Config) {
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
				log.Printf("Server is starting at :%d", config.ServerPort)
				http.ListenAndServe(fmt.Sprintf(":%d", config.ServerPort), handler)
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Server stopping...")
			return nil
		},
	})
}
