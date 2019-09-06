package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Post -
type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var posts []Post

// GetPost - gytftfu
func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range posts {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
		return
	}
	json.NewEncoder(w).Encode(&Post{})
}
