package controllers

import (
	"encoding/json"
	"mybookstore/pkg/config"
	"mybookstore/pkg/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var existing models.Author
	if err := config.DB.Where("name = ?", author.Name).First(&existing).Error; err == nil {
		http.Error(w, "Author already exists", http.StatusConflict)
		return
	}

	if err := config.DB.Create(&author).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(author)
}

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	var authors []models.Author
	result := config.DB.Find(&authors)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authors)
}

func GetAuthorByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid author ID", http.StatusBadRequest)
		return
	}
	var author models.Author
	result := config.DB.First(&author, id)
	if result.Error != nil {
		http.Error(w, "Author not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(author)
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid author ID", http.StatusBadRequest)
		return
	}
	var author models.Author
	result := config.DB.First(&author, id)
	if result.Error != nil {

		http.Error(w, "Author not found", http.StatusNotFound)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	config.DB.Save(&author)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(author)
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid author ID", http.StatusBadRequest)
		return
	}
	var author models.Author
	result := config.DB.First(&author, id)
	if result.Error != nil {
		http.Error(w, "Author not found", http.StatusNotFound)
		return
	}
	config.DB.Delete(&author)
	w.WriteHeader(http.StatusNoContent)
}
