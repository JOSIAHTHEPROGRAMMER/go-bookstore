package controllers

import (
	"encoding/json"
	"mybookstore/pkg/config"
	"mybookstore/pkg/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	result := config.DB.Find(&books)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var existing models.Book
	if err := config.DB.Where("title = ?", book.Title).First(&existing).Error; err == nil {
		http.Error(w, "Book already exists", http.StatusConflict)
		return
	}

	if err := config.DB.Create(&book).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	var book models.Book
	result := config.DB.First(&book, id)
	if result.Error != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	var book models.Book
	result := config.DB.First(&book, id)
	if result.Error != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	config.DB.Save(&book)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var book models.Book
	result := config.DB.First(&book, id)
	if result.Error != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	config.DB.Delete(&book)
	w.WriteHeader(http.StatusNoContent)
}

func SearchBooks(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	author := queryParams.Get("author")
	title := queryParams.Get("title")
	var books []models.Book
	dbQuery := config.DB
	if author != "" {
		dbQuery = dbQuery.Where("author LIKE ?", "%"+author+"%")
	}
	if title != "" {
		dbQuery = dbQuery.Where("title LIKE ?", "%"+title+"%")
	}
	result := dbQuery.Find(&books)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)

		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
