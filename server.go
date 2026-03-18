package main

import (
	"encoding/json"
	"fmt"
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
	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&Allpost)
}

func addPostById(w http.ResponseWriter, r *http.Request, splitstring []string) {
	idval, err := strconv.Atoi(splitstring[2])
	if err != nil {
		log.Fatal("Invalid integer value")
		return
	}
	if HandleGetError(w, r) {
		return
	}
	for i, data := range Allpost {
		if i == idval {
			w.Header().Set("Content-Type", "aplication/json")
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(data)
			break
		} else {
			HandleNotFoundError(w, r)
		}
	}
}

func fetchPost(w http.ResponseWriter, r *http.Request, stringval []string) {
	idval, err := strconv.Atoi(stringval[2])
	if HandleStringConvError(err) {
		return
	}
	for i, j := range Allpost {
		if i == idval {
			w.Header().Set("Content-Type", "aplication/json")
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(j)
		} else {
			HandleNotFoundError(w, r)
		}
	}

}
func addCommentById(w http.ResponseWriter, r *http.Request, stringval []string) {
	idval, err := strconv.Atoi(stringval[2])
	if HandleStringConvError(err) {
		return
	}
	var post postData
	json.NewDecoder(r.Body).Decode(post)
	Allpost[idval] = append(Allpost[idval], postData{
		Id:     post.Id,
		PostId: post.PostId,
		Text:   post.Text,
	})
	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusAccepted)
}
func getPostById(w http.ResponseWriter, r *http.Request, id string) {
	idval, err := strconv.Atoi(id)
	if HandleStringConvError(err) {
		return
	}
	for i, j := range Allpost {
		if i == idval {
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(j)
		} else {
			HandleNotFoundError(w, r)
		}
	}
}
func detelePostById(w http.ResponseWriter, r *http.Request, splitstr []string) {
	idval, err := strconv.Atoi(splitstr[2])
	if HandleStringConvError(err) {
		return
	}
	for i, _ := range Allpost {
		if i == idval {
			delete(Allpost, i)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(Allpost)
		} else {
			HandleNotFoundError(w, r)
		}
	}
}
func handlePostFunc(w http.ResponseWriter, r *http.Request) {
	stringval := strings.TrimPrefix(r.URL.Path, "/posts")
	splitstring := strings.Split(stringval, "/")
	fmt.Println("****", len(splitstring), stringval)
	switch r.Method {

	case http.MethodGet:
		// get post/{id}/comment
		if len(splitstring) == 4 && splitstring[3] == "comment" {
			fetchPost(w, r, splitstring)
		}
		// get /post/{id}
		if len(splitstring) == 3 {
			getPostById(w, r, splitstring[2])
		}
	case http.MethodPost:
		//post /post/{id}/comment
		if len(splitstring) == 4 && splitstring[3] == "comment" {
			addPostById(w, r, splitstring)
		}
		//post /post/{id}
		if len(splitstring) == 3 {
			addCommentById(w, r, splitstring)
		}
	case http.MethodDelete:
		//delete /post/{id}
		if len(splitstring) == 2 {
			detelePostById(w, r, splitstring)
		}

	default:
		http.NotFound(w, r)
	}

}
