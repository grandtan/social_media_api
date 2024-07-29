package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"social_media_app/database"
	"social_media_app/models"
	"social_media_app/routes"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *mux.Router {
	router := mux.NewRouter()
	routes.RegisterRoutes(router)
	return router
}

func TestCreatePost(t *testing.T) {
	database.Connect()
	router := setupRouter()

	post := models.Post{
		UserID:  1,
		Content: "Test Content",
	}
	jsonPost, _ := json.Marshal(post)

	req, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonPost))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	var createdPost models.Post
	json.Unmarshal(resp.Body.Bytes(), &createdPost)
	assert.Equal(t, post.Content, createdPost.Content)
}

func TestGetPost(t *testing.T) {
	database.Connect()
	router := setupRouter()

	// Create a post first
	post := models.Post{
		UserID:  1,
		Content: "Test Content",
	}
	database.DB.Create(&post)

	req, _ := http.NewRequest("GET", fmt.Sprintf("/posts/%d", post.ID), nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	var fetchedPost models.Post
	json.Unmarshal(resp.Body.Bytes(), &fetchedPost)
	assert.Equal(t, post.Content, fetchedPost.Content)
}

func TestDeletePost(t *testing.T) {
	database.Connect()
	router := setupRouter()

	// Create a post first
	post := models.Post{
		UserID:  1,
		Content: "Test Content",
	}
	database.DB.Create(&post)

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/posts/%d", post.ID), nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}
