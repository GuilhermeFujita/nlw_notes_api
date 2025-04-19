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
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", config.ServerPort),
		Handler: cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
			AllowedHeaders: []string{"Content-Type"},
		}).Handler(router),
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Printf("Server is starting at %s", srv.Addr)
			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatalf("ListenAndServe error %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Stopping server...")
			return srv.Shutdown(ctx)
		},
	})
}
