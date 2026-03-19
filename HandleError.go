package main

import (
	"encoding/json"
	"log"
	"net/http"
)

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

func HandleStringConvError(err error) bool {
	if err != nil {
		log.Fatal("Error in conversrion")
		return true
	} else {
		return false
	}
}

func HandleNotFoundError(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"error": "Not Found",
	})

}

func HandlePostPresent(id int, maptoCheck map[int]postData) bool {
	_, ok := maptoCheck[id]
	if !ok {
		return false
	} else {
		return true
	}
}

func HandleCommentPresent(id int, maptoCheck map[int][]Comment) bool {
	_, ok := maptoCheck[id]
	if !ok {
		return false
	} else {
		return true
	}
}
