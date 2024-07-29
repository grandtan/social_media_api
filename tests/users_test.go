package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"social_media_app/database"
	"social_media_app/models"
	"social_media_app/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupUsers() {
	database.Connect()
	database.DB.Exec("DELETE FROM users")
	database.DB.Exec("DELETE FROM sqlite_sequence WHERE name='users'")
	user := models.User{Name: "Test User", Email: "test@example.com"}
	database.DB.Create(&user)
	token, _ := utils.GenerateJWT(user.ID)
	testToken = token
}

func teardownUsers() {
	database.DB.Exec("DELETE FROM users")
	database.DB.Exec("DELETE FROM sqlite_sequence WHERE name='users'")
}

func TestCreateUser(t *testing.T) {
	setupUsers()
	defer teardownUsers()

	router := SetupRouter()

	user := models.User{
		Name:  "New User",
		Email: "new@example.com",
	}
	jsonUser, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonUser))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+testToken)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	fmt.Println("Response Body:", resp.Body.String())

	assert.Equal(t, http.StatusOK, resp.Code)
	var createdUser models.User
	json.Unmarshal(resp.Body.Bytes(), &createdUser)
	assert.Equal(t, user.Email, createdUser.Email)
}

func TestGetUser(t *testing.T) {
	setupUsers()
	defer teardownUsers()

	router := SetupRouter()

	// ลบผู้ใช้เพื่อให้แน่ใจว่าไม่มีผู้ใช้ที่ซ้ำกัน
	database.DB.Exec("DELETE FROM users")
	user := models.User{
		Name:  "Test User",
		Email: "test@example.com",
	}
	if err := database.DB.Create(&user).Error; err != nil {
		t.Fatalf("could not create user: %v", err)
	}

	req, _ := http.NewRequest("GET", fmt.Sprintf("/users/%d", user.ID), nil)
	req.Header.Set("Authorization", "Bearer "+testToken)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	fmt.Println("Response Body:", resp.Body.String())

	assert.Equal(t, http.StatusOK, resp.Code)
	var fetchedUser models.User
	json.Unmarshal(resp.Body.Bytes(), &fetchedUser)
	assert.Equal(t, user.Email, fetchedUser.Email)
}

func TestDeleteUser(t *testing.T) {
	setupUsers()
	defer teardownUsers()

	router := SetupRouter()

	database.DB.Exec("DELETE FROM users")
	user := models.User{
		Name:  "Test User",
		Email: "test@example.com",
	}
	if err := database.DB.Create(&user).Error; err != nil {
		t.Fatalf("could not create user: %v", err)
	}

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/users/%d", user.ID), nil)
	req.Header.Set("Authorization", "Bearer "+testToken)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	fmt.Println("Response Body:", resp.Body.String())

	assert.Equal(t, http.StatusOK, resp.Code)
}
