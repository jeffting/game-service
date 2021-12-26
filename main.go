package main

import "github.com/jeffting/game-service/api"

func main() {
	server := api.NewServer()
	api.Routes(server)
}
