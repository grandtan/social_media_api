package controllers

import (
	"encoding/json"
	"net/http"
	"social_media_app/database"
	"social_media_app/models"

	"github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	json.NewDecoder(r.Body).Decode(&post)
	database.DB.Create(&post)
	json.NewEncoder(w).Encode(post)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var post models.Post
	database.DB.First(&post, params["id"])
	json.NewEncoder(w).Encode(post)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post
	database.DB.Find(&posts)
	json.NewEncoder(w).Encode(posts)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var post models.Post
	database.DB.First(&post, params["id"])
	json.NewDecoder(r.Body).Decode(&post)
	database.DB.Save(&post)
	json.NewEncoder(w).Encode(post)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var post models.Post
	database.DB.Delete(&post, params["id"])
	json.NewEncoder(w).Encode("Post deleted")
}
