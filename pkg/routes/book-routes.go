package routes

import (
	"mybookstore/pkg/controllers"

	"github.com/gorilla/mux"
)

func BookRoutes(router *mux.Router) {
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBookByID).Methods("GET")

	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")

	router.HandleFunc("/books/search", controllers.SearchBooks).Methods("GET")
}
