package controllers

import (
	"net/http"
	"social_media_app/database"
	"social_media_app/models"
	"social_media_app/utils"

	"github.com/gin-gonic/gin"
)

// CreateUser handles the creation of a new user
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid input")
		return
	}

	if user.Name == "" || user.Email == "" {
		utils.RespondError(c, http.StatusBadRequest, "Name and Email are required")
		return
	}

	if err := database.DB.Create(&user).Error; err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusOK, user)
}

// Login handles user authentication
func Login(c *gin.Context) {
	var requestUser models.User
	if err := c.ShouldBindJSON(&requestUser); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid input")
		return
	}

	var user models.User
	if err := database.DB.Where("email = ?", requestUser.Email).First(&user).Error; err != nil {
		utils.RespondError(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	// Here you should check the password. This is a simple example, you should hash and compare passwords.
	if requestUser.Name != user.Name {
		utils.RespondError(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Could not generate token")
		return
	}

	utils.RespondJSON(c, http.StatusOK, gin.H{"token": token})
}

// GetUser handles fetching a user by ID
func GetUser(c *gin.Context) {
	var user models.User
	if err := database.DB.First(&user, c.Param("id")).Error; err != nil {
		utils.RespondError(c, http.StatusNotFound, "User not found")
		return
	}
	utils.RespondJSON(c, http.StatusOK, user)
}

// GetUsers handles fetching all users
func GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	utils.RespondJSON(c, http.StatusOK, users)
}

// UpdateUser handles updating a user by ID
func UpdateUser(c *gin.Context) {
	var user models.User
	if err := database.DB.First(&user, c.Param("id")).Error; err != nil {
		utils.RespondError(c, http.StatusNotFound, "User not found")
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid input")
		return
	}

	database.DB.Save(&user)
	utils.RespondJSON(c, http.StatusOK, user)
}

// DeleteUser handles deleting a user by ID
func DeleteUser(c *gin.Context) {
	var user models.User
	if err := database.DB.Delete(&user, c.Param("id")).Error; err != nil {
		utils.RespondError(c, http.StatusNotFound, "User not found")
		return
	}
	utils.RespondJSON(c, http.StatusOK, "User deleted")
}
