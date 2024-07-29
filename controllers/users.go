package controllers

import (
	"encoding/json"
	"net/http"
	"social_media_app/database"
	"social_media_app/models"
	"social_media_app/utils"

	"github.com/gorilla/mux"
)

// CreateUser handles the creation of a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid input")
		return
	}

	if user.Name == "" || user.Email == "" {
		utils.RespondError(w, http.StatusBadRequest, "Name and Email are required")
		return
	}

	if err := database.DB.Create(&user).Error; err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, user)
}

// GetUser handles fetching a user by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User
	if err := database.DB.First(&user, params["id"]).Error; err != nil {
		utils.RespondError(w, http.StatusNotFound, "User not found")
		return
	}
	utils.RespondJSON(w, http.StatusOK, user)
}

// GetUsers handles fetching all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	database.DB.Find(&users)
	utils.RespondJSON(w, http.StatusOK, users)
}

// UpdateUser handles updating a user by ID
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User
	if err := database.DB.First(&user, params["id"]).Error; err != nil {
		utils.RespondError(w, http.StatusNotFound, "User not found")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid input")
		return
	}

	database.DB.Save(&user)
	utils.RespondJSON(w, http.StatusOK, user)
}

// DeleteUser handles deleting a user by ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User
	if err := database.DB.Delete(&user, params["id"]).Error; err != nil {
		utils.RespondError(w, http.StatusNotFound, "User not found")
		return
	}
	utils.RespondJSON(w, http.StatusOK, "User deleted")
}
