package api

import (
	"fmt"
	"log"
	"net/http"

	v1Users "github.com/jeffting/game-service/api/v1/users"

	"github.com/gorilla/mux"
)

const port = "8083"

func Routes(server *Server) {
	router := mux.NewRouter().StrictSlash(true)
	fmt.Println("Running on localhost:" + port)

	router.HandleFunc("/v1/users/{username}", v1Users.GetUser(server.clients)).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
