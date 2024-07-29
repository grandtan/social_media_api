package controllers

import (
	"encoding/json"
	"net/http"
	"social_media_app/database"
	"social_media_app/models"

	"github.com/gorilla/mux"
)

// CreatePost handles the creation of a new post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if post.UserID == 0 || post.Content == "" {
		http.Error(w, "UserID and Content are required", http.StatusBadRequest)
		return
	}

	if err := database.DB.Create(&post).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(post)
}

// GetPost handles fetching a post by ID
func GetPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var post models.Post
	if err := database.DB.First(&post, params["id"]).Error; err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(post)
}

// GetPosts handles fetching all posts
func GetPosts(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post
	database.DB.Find(&posts)
	json.NewEncoder(w).Encode(posts)
}

// UpdatePost handles updating a post by ID
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var post models.Post
	if err := database.DB.First(&post, params["id"]).Error; err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	database.DB.Save(&post)
	json.NewEncoder(w).Encode(post)
}

// DeletePost handles deleting a post by ID
func DeletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var post models.Post
	if err := database.DB.Delete(&post, params["id"]).Error; err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode("Post deleted")
}
