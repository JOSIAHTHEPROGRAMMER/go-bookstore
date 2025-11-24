package routes

import (
	"mybookstore/pkg/controllers"

	"github.com/gorilla/mux"
)

func AuthorRoutes(router *mux.Router) {

	router.HandleFunc("/authors", controllers.GetAuthors).Methods("GET")
	router.HandleFunc("/authors/{id}", controllers.GetAuthorByID).Methods("GET")
	router.HandleFunc("/authors", controllers.CreateAuthor).Methods("POST")
	router.HandleFunc("/authors/{id}", controllers.UpdateAuthor).Methods("PUT")
	router.HandleFunc("/authors/{id}", controllers.DeleteAuthor).Methods("DELETE")
}
