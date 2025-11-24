package routes

import (
	"mybookstore/pkg/controllers"

	"github.com/gorilla/mux"
)

func PublisherRoutes(router *mux.Router) {

	router.HandleFunc("/publishers", controllers.GetPublishers).Methods("GET")
	router.HandleFunc("/publishers/{id}", controllers.GetPublisherByID).Methods("GET")
	router.HandleFunc("/publishers", controllers.CreatePublisher).Methods("POST")
	router.HandleFunc("/publishers/{id}", controllers.UpdatePublisher).Methods("PUT")
	router.HandleFunc("/publishers/{id}", controllers.DeletePublisher).Methods("DELETE")
}
