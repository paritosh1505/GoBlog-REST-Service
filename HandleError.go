package main

import "net/http"

func HandleGetError(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodGet {
		http.Error(w, "ONLY GET METHOD ALLOWED", http.StatusMethodNotAllowed)
		return true
	} else {
		return false
	}
}
func HandlePostError(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodPost {
		http.Error(w, "ONLY POST METHOD ALLOWED", http.StatusMethodNotAllowed)
		return true
	} else {
		return false
	}
}
