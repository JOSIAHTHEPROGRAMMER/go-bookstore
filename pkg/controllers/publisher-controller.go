package controllers

import (
	"encoding/json"
	"mybookstore/pkg/config"
	"mybookstore/pkg/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetPublishers(w http.ResponseWriter, r *http.Request) {
	var publishers []models.Publisher
	result := config.DB.Find(&publishers)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(publishers)
}

func GetPublisherByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid publisher ID", http.StatusBadRequest)
		return
	}
	var publisher models.Publisher
	result := config.DB.First(&publisher, id)
	if result.Error != nil {
		http.Error(w, "Publisher not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(publisher)
}

func CreatePublisher(w http.ResponseWriter, r *http.Request) {
	var publisher models.Publisher
	err := json.NewDecoder(r.Body).Decode(&publisher)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var existing models.Publisher
	if err := config.DB.Where("name = ?", publisher.Name).First(&existing).Error; err == nil {
		http.Error(w, "Publisher already exists", http.StatusConflict)
		return
	}

	result := config.DB.Create(&publisher)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(publisher)
}

func UpdatePublisher(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid publisher ID", http.StatusBadRequest)
		return
	}
	var publisher models.Publisher
	result := config.DB.First(&publisher, id)
	if result.Error != nil {
		http.Error(w, "Publisher not found", http.StatusNotFound)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&publisher)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result = config.DB.Save(&publisher)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(publisher)
}

func DeletePublisher(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid publisher ID", http.StatusBadRequest)
		return
	}
	var publisher models.Publisher
	result := config.DB.Delete(&publisher, id)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
