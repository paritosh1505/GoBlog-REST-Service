package main

import "net/http"

func main() {
	http.HandleFunc("/post", createPost)
	http.HandleFunc("/getPost", getPost)
	http.HandleFunc("/post/{id}/comment", addComment)
	http.HandleFunc("/getPost/{id}", fetchPost)
	http.HandleFunc("/post/{id}", AddPostById)
	http.ListenAndServe(":8080", nil)
}
