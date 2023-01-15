package routes

import (
	"erlangga-final-task/handlers"
	"erlangga-final-task/pkg/middleware"
	"erlangga-final-task/pkg/mysql"
	"erlangga-final-task/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	channelRepository := repositories.RepositoryChannel(mysql.DB)
	h := handlers.HandlerAuth(channelRepository)

	r.HandleFunc("/register", h.Register).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST")
	r.HandleFunc("/check-auth", middleware.Auth(h.CheckAuth)).Methods("GET")
}
