package routes

import (
	"erlangga-final-task/handlers"
	"erlangga-final-task/pkg/middleware"
	"erlangga-final-task/pkg/mysql"
	"erlangga-final-task/repositories"

	"github.com/gorilla/mux"
)

func ChannelRoutes(r *mux.Router) {
	channelRepository := repositories.RepositoryChannel(mysql.DB)
	h := handlers.HandlerChannel(channelRepository)

	r.HandleFunc("/channels", h.FindChannels).Methods("GET")
	r.HandleFunc("/channel/{id}", h.GetChannel).Methods("GET")
	r.HandleFunc("/channel/{id}", middleware.Auth(middleware.UploadCover(middleware.UploadPhoto(h.UpdateChannel)))).Methods("PATCH")
	r.HandleFunc("/channel/{id}", middleware.Auth(h.DeleteChannel)).Methods("DELETE")

	r.HandleFunc("/plusSubs/{id}", middleware.Auth(h.PlusSubscriber)).Methods("PATCH")
	r.HandleFunc("/minusSubs/{id}", middleware.Auth(h.MinusSubscriber)).Methods("PATCH")
}
