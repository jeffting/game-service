package api

import (
	"github.com/jeffting/game-service/clients"
	"github.com/jeffting/game-service/clients/db"
)

type Server struct {
	clients clients.Clients
}

func NewServer() *Server {
	gameDao := db.NewGameClient()
	clients := clients.Clients{
		DB: gameDao,
	}
	return &Server{
		clients: clients,
	}
}
