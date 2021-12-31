package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/willcliffy/keydream-server/common"
	game_models "github.com/willcliffy/keydream-server/world/models"
	"github.com/willcliffy/keydream-server/lobby"
)

func main() {
	lobbyHandler := lobby.LobbyHandler{
		Worlds: make(map[common.WorldID]game_models.WorldBroadcast),
	}

	server := http.Server{
		Addr:    "0.0.0.0:80",
		Handler: ConnectRouter(lobbyHandler),
	}

	shutdown := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)

		<-sigint

		if err := server.Shutdown(context.Background()); err != nil {
			log.Printf("shutdown error: %v\n", err)
		}
		close(shutdown)
	}()

	go lobbyHandler.ControlLoop()

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("err: %v\n", err)
	}

	<-shutdown
}

func ConnectRouter(
	lobbyHandler lobby.LobbyHandler,
) chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.StripSlashes)
	router.Use(middleware.Recoverer)

	router.Route("/api/v1", func(r chi.Router) {
		r.Post("/connect", lobbyHandler.ConnectHandler)
		r.Post("/join", lobbyHandler.JoinHandler)

		r.Post("/worlds", lobbyHandler.UpdateWorldHandler)
	})

	return router
}
