package main

import (
	"github.com/alvinatthariq/go-restfull/handler"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// router.HandleFunc("/posts", handler.GetPosts).Methods("GET")
	// router.HandleFunc("/posts", createPost).Methods("POST")
	router.HandleFunc("/posts/{id}", handler.GetPost).Methods("GET")
	// router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	// router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")
}
