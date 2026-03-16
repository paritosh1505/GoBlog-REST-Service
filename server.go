package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type postData struct {
	Id     int    `json:"id"`
	PostId int    `json:"post_id"`
	Text   string `json:"text"`
}

var Allpost = make(map[int][]postData)

func createPost(w http.ResponseWriter, r *http.Request) {
	if HandlePostError(w, r) {
		return
	}
	var post postData
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}
	post.Id += 1
	if post.PostId == 0 {
		post.PostId = post.Id
	}
	Allpost[len(Allpost)] = append(Allpost[len(Allpost)], post)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(Allpost)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	if HandleGetError(w, r) {
		return
	}

	json.NewEncoder(w).Encode(&Allpost)
}
func addComment(w http.ResponseWriter, r *http.Request) {
	val := r.URL.Path
	var post postData
	commentId, err := strconv.Atoi(strings.Split(val, "/")[2])
	if err != nil {
		log.Fatal("Invalid string conversion")
		return
	}
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Invali json type", http.StatusBadRequest)
		return
	}
	for i, postitr := range Allpost {
		if i == commentId {
			postid := postitr[len(postitr)-1].PostId
			Allpost[i] = append(Allpost[i], postData{
				Id:     postitr[i].Id,
				PostId: postid + 1,
				Text:   post.Text,
			})
		}

	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&Allpost)
}

func fetchPost(w http.ResponseWriter, r *http.Request) {
	idstring := r.URL.Path
	idval, err := strconv.Atoi(strings.Split(idstring, "/")[2])
	if err != nil {
		http.Error(w, "invalid get id", http.StatusBadRequest)
	}
	if HandleGetError(w, r) {
		return
	}

	for i, data := range Allpost {
		if i == idval {
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(data)
			break
		}
	}
}

func AddPostById(w http.ResponseWriter, r *http.Request) {
	if HandlePostError(w, r) {
		return
	}
}
