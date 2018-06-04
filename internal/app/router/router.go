package router

import (
	"github.com/gorilla/mux"

	"github.com/austinpgraham/chocolate.server/internal/app/user"
)

func DefineRoutes(router *mux.Router) {
	router.HandleFunc("/users", user.CreateUser).Methods("POST")
}
