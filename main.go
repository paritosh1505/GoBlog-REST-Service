package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/post", handlePostFunc)
	http.HandleFunc("/post/", handlePostFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
