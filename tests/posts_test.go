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

func setupPosts() {
	database.Connect()
	database.DB.Exec("DELETE FROM posts")
	database.DB.Exec("DELETE FROM sqlite_sequence WHERE name='posts'")
	user := models.User{Name: "Test User", Email: "test@example.com"}
	database.DB.Create(&user)
	token, _ := utils.GenerateJWT(user.ID)
	testToken = token
}

func teardownPosts() {
	database.DB.Exec("DELETE FROM posts")
	database.DB.Exec("DELETE FROM sqlite_sequence WHERE name='posts'")
}

func TestCreatePost(t *testing.T) {
	setupPosts()
	defer teardownPosts()

	router := SetupRouter()

	post := models.Post{
		UserID:  1,
		Content: "Test Content",
	}
	jsonPost, _ := json.Marshal(post)

	req, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonPost))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+testToken)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	fmt.Println("Response Body:", resp.Body.String())

	assert.Equal(t, http.StatusOK, resp.Code)
	var createdPost models.Post
	json.Unmarshal(resp.Body.Bytes(), &createdPost)
	assert.Equal(t, post.Content, createdPost.Content)
}

func TestGetPost(t *testing.T) {
	setupPosts()
	defer teardownPosts()

	router := SetupRouter()

	post := models.Post{
		UserID:  1,
		Content: "Test Content",
	}
	if err := database.DB.Create(&post).Error; err != nil {
		t.Fatalf("could not create post: %v", err)
	}

	req, _ := http.NewRequest("GET", fmt.Sprintf("/posts/%d", post.ID), nil)
	req.Header.Set("Authorization", "Bearer "+testToken)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	fmt.Println("Response Body:", resp.Body.String())

	assert.Equal(t, http.StatusOK, resp.Code)
	var fetchedPost models.Post
	json.Unmarshal(resp.Body.Bytes(), &fetchedPost)
	assert.Equal(t, post.Content, fetchedPost.Content)
}

func TestDeletePost(t *testing.T) {
	setupPosts()
	defer teardownPosts()

	router := SetupRouter()

	post := models.Post{
		UserID:  1,
		Content: "Test Content",
	}
	if err := database.DB.Create(&post).Error; err != nil {
		t.Fatalf("could not create post: %v", err)
	}

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/posts/%d", post.ID), nil)
	req.Header.Set("Authorization", "Bearer "+testToken)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	fmt.Println("Response Body:", resp.Body.String())

	assert.Equal(t, http.StatusOK, resp.Code)
}
