package router

import (
	"github.com/gorilla/mux"

	"github.com/austinpgraham/chocolate.server/internal/app/course"
	"github.com/austinpgraham/chocolate.server/internal/app/neighborhood"
	"github.com/austinpgraham/chocolate.server/internal/app/user"
)

func DefineRoutes(router *mux.Router) {
	router.HandleFunc("/users", user.CreateUser).Methods("POST")
	router.HandleFunc("/auth/login", user.Login).Methods("POST")
	router.HandleFunc("/users/{username}", user.GetUser).Methods("GET")
	router.HandleFunc("/auth/logout", user.Logout).Methods("GET")
	router.HandleFunc("/chocolate/courses", course.CreateCourse).Methods("POST")
	router.HandleFunc("/doorbell/groups", neighborhood.CreateNeighborhood).Methods("POST")
	router.HandleFunc("/doorbell/groups/{neighborhood}", neighborhood.GetNeighborhood).Methods("GET")
	router.HandleFunc("/doorbell/groups", neighborhood.GetOwnedNeighborhoods).Methods("GET")
}
