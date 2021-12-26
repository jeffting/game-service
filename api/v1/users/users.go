package users

import (
	"context"
	"log"
	"net/http"

	"github.com/jeffting/game-service/clients"
	"github.com/jeffting/game-service/utils"

	"github.com/gorilla/mux"
)

func GetUser(clients clients.Clients) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		vars := mux.Vars(r)
		username := vars["username"]
		idk, err := clients.DB.GetUserByUsername(ctx, username)
		if err != nil {
			log.Print("test", err)
			utils.SendError(w, http.StatusInternalServerError, err)
		}
		utils.SendJSON(w, idk)
	}
}
