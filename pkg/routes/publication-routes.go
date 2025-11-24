package routes

import (
	"mybookstore/pkg/controllers"

	"github.com/gorilla/mux"
)

func PublicationRoutes(router *mux.Router) {
	router.HandleFunc("/publications", controllers.GetPublications).Methods("GET")
	router.HandleFunc("/publications/{id}", controllers.GetPublicationByID).Methods("GET")
	router.HandleFunc("/publications", controllers.CreatePublication).Methods("POST")
	router.HandleFunc("/publications/{id}", controllers.UpdatePublication).Methods("PUT")
	router.HandleFunc("/publications/{id}", controllers.DeletePublication).Methods("DELETE")
	router.HandleFunc("/publications/search", controllers.SearchPublications).Methods("GET")
}
