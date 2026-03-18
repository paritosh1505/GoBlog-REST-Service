package main

import (
	"log"
	"net/http"
)

func main() {
	//http.HandleFunc("/post", createPost)
	//http.HandleFunc("/post", getPost)
	//http.HandleFunc("/post/{id}/comment", addComment)
	//http.HandleFunc("/post/{id}", fetchPost)
	//http.HandleFunc("/post/{id}", AddPostById)
	http.HandleFunc("/post/", handlePostFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
