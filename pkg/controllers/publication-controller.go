package controllers

import (
	"encoding/json"
	"mybookstore/pkg/config"
	"mybookstore/pkg/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetPublications(w http.ResponseWriter, r *http.Request) {
	var publications []models.Publication
	result := config.DB.Find(&publications)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(publications)
}

func GetPublicationByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid publication ID", http.StatusBadRequest)
		return
	}
	var publication models.Publication
	result := config.DB.First(&publication, id)
	if result.Error != nil {
		http.Error(w, "Publication not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(publication)
}

func CreatePublication(w http.ResponseWriter, r *http.Request) {
	var publication models.Publication
	if err := json.NewDecoder(r.Body).Decode(&publication); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var book models.Book
	if err := config.DB.First(&book, publication.BookID).Error; err != nil {
		http.Error(w, "Book not found", http.StatusBadRequest)
		return
	}
	var author models.Author
	if err := config.DB.First(&author, publication.AuthorID).Error; err != nil {
		http.Error(w, "Author not found", http.StatusBadRequest)
		return
	}
	var publisher models.Publisher
	if err := config.DB.First(&publisher, publication.PublisherID).Error; err != nil {
		http.Error(w, "Publisher not found", http.StatusBadRequest)
		return
	}

	var existing models.Publication
	if err := config.DB.
		Where("book_id = ? AND author_id = ? AND publisher_id = ?", publication.BookID, publication.AuthorID, publication.PublisherID).
		First(&existing).Error; err == nil {
		http.Error(w, "Publication already exists", http.StatusConflict)
		return
	}

	if err := config.DB.Create(&publication).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(publication)
}

func UpdatePublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid publication ID", http.StatusBadRequest)
		return
	}
	var publication models.Publication
	result := config.DB.First(&publication, id)
	if result.Error != nil {
		http.Error(w, "Publication not found", http.StatusNotFound)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&publication)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result = config.DB.Save(&publication)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(publication)
}

func DeletePublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid publication ID", http.StatusBadRequest)
		return
	}
	var publication models.Publication
	result := config.DB.Delete(&publication, id)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func SearchPublications(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	title := queryParams.Get("title")
	author := queryParams.Get("author")
	var publications []models.Publication
	query := config.DB.Model(&models.Publication{})
	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}
	if author != "" {
		query = query.Where("author LIKE ?", "%"+author+"%")
	}
	result := query.Find(&publications)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(publications)
}
