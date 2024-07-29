package controllers

import (
	"encoding/json"
	"net/http"
	"social_media_app/database"
	"social_media_app/models"
	"social_media_app/utils"

	"github.com/gorilla/mux"
)

// CreatePost handles the creation of a new post
func CreatePost(w http.ResponseWriter, r *http.Request) {
    var post models.Post
    if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
        utils.RespondError(w, http.StatusBadRequest, "Invalid input")
        return
    }

    if post.UserID == 0 || post.Content == "" {
        utils.RespondError(w, http.StatusBadRequest, "UserID and Content are required")
        return
    }

    if err := database.DB.Create(&post).Error; err != nil {
        utils.RespondError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondJSON(w, http.StatusOK, post)
}

// GetPost handles fetching a post by ID
func GetPost(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var post models.Post
    if err := database.DB.First(&post, params["id"]).Error; err != nil {
        utils.RespondError(w, http.StatusNotFound, "Post not found")
        return
    }
    utils.RespondJSON(w, http.StatusOK, post)
}

// GetPosts handles fetching all posts
func GetPosts(w http.ResponseWriter, r *http.Request) {
    var posts []models.Post
    database.DB.Find(&posts)
    utils.RespondJSON(w, http.StatusOK, posts)
}

// UpdatePost handles updating a post by ID
func UpdatePost(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var post models.Post
    if err := database.DB.First(&post, params["id"]).Error; err != nil {
        utils.RespondError(w, http.StatusNotFound, "Post not found")
        return
    }

    if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
        utils.RespondError(w, http.StatusBadRequest, "Invalid input")
        return
    }

    database.DB.Save(&post)
    utils.RespondJSON(w, http.StatusOK, post)
}

// DeletePost handles deleting a post by ID
func DeletePost(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var post models.Post
    if err := database.DB.Delete(&post, params["id"]).Error; err != nil {
        utils.RespondError(w, http.StatusNotFound, "Post not found")
        return
    }
    utils.RespondJSON(w, http.StatusOK, "Post deleted")
}
