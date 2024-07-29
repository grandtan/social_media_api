package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"social_media_app/database"
	"social_media_app/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	database.Connect()
	router := SetupRouter()

	user := models.User{
		Name:  "Test User",
		Email: "test@example.com",
	}
	jsonUser, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonUser))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	var createdUser models.User
	json.Unmarshal(resp.Body.Bytes(), &createdUser)
	assert.Equal(t, user.Email, createdUser.Email)
}

func TestGetUser(t *testing.T) {
	database.Connect()
	router := SetupRouter()

	// Create a user first
	user := models.User{
		Name:  "Test User",
		Email: "test@example.com",
	}
	if err := database.DB.Create(&user).Error; err != nil {
		t.Fatalf("could not create user: %v", err)
	}

	req, _ := http.NewRequest("GET", fmt.Sprintf("/users/%d", user.ID), nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	var fetchedUser models.User
	json.Unmarshal(resp.Body.Bytes(), &fetchedUser)
	assert.Equal(t, user.Email, fetchedUser.Email)
}

func TestDeleteUser(t *testing.T) {
	database.Connect()
	router := SetupRouter()

	// Create a user first
	user := models.User{
		Name:  "Test User",
		Email: "test@example.com",
	}
	if err := database.DB.Create(&user).Error; err != nil {
		t.Fatalf("could not create user: %v", err)
	}

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/users/%d", user.ID), nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}
