package routes

import (
	"erlangga-final-task/handlers"
	"erlangga-final-task/pkg/middleware"
	"erlangga-final-task/pkg/mysql"
	"erlangga-final-task/repositories"

	"github.com/gorilla/mux"
)

func SubscribeRoutes(r *mux.Router) {
	subscribeRepository := repositories.RepositorySubscribe(mysql.DB)
	h := handlers.HandlerSubscribe(subscribeRepository)

	r.HandleFunc("/subscribes", middleware.Auth(h.FindSubscribes)).Methods("GET")
	r.HandleFunc("/subscribe/{id}", middleware.Auth(h.GetSubscribe)).Methods("GET")

	r.HandleFunc("/subscribeByOther/{id}", middleware.Auth(h.GetSubscribeByOther)).Methods("GET")

	r.HandleFunc("/subscribe/{id}", middleware.Auth(h.CreateSubscribe)).Methods("POST")
	r.HandleFunc("/subscribe", middleware.Auth(h.DeleteSubscribe)).Methods("DELETE")

	r.HandleFunc("/subscription", middleware.Auth(h.GetSubscription)).Methods("GET")

}
