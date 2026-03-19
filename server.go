package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type postData struct {
	Id      int    `json:"Id"`
	Post_id int    `json:"postId"`
	Caption string `json:"text"`
}
type Comment struct {
	Id       int    `json:"commentId"`
	Postid   int    `json:"postId"`
	CommPost string `json:"addedComment"`
}

var Addpost = make(map[int]postData)
var AddComment = make(map[int][]Comment)
var key = 0

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
	key = len(Addpost) + 1
	Addpost[key] = postData{
		Caption: post.Caption,
		Id:      len(Addpost),
		Post_id: post.Id + 1,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Addpost)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	if HandleGetError(w, r) {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Addpost)

}
func addComment(w http.ResponseWriter, r *http.Request) {
	val := r.URL.Path
	var post postData
	_, err := strconv.Atoi(strings.Split(val, "/")[2])
	if err != nil {
		log.Fatal("Invalid string conversion")
		return
	}
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Invali json type", http.StatusBadRequest)
		return
	}

}

func addCommentById(w http.ResponseWriter, r *http.Request, splitstring []string) {
	idval, err := strconv.Atoi(splitstring[2])
	if err != nil {
		log.Fatal("Invalid integer value")
		return
	}
	if HandlePostError(w, r) {
		return
	}
	var comment Comment
	status := HandlePostPresent(idval, Addpost)
	if status {
		json.NewDecoder(r.Body).Decode(&comment)
		postid := len(AddComment)
		AddComment[postid] = append(AddComment[postid], Comment{
			Id:       idval,
			Postid:   postid,
			CommPost: comment.CommPost,
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(AddComment)
	} else {
		HandleNotFoundError(w, r)
	}

}

func fetchCommentById(w http.ResponseWriter, r *http.Request, stringval []string) {
	idval, err := strconv.Atoi(stringval[2])
	if HandleStringConvError(err) {
		return
	}
	var fetchComment []Comment
	if HandleCommentPresent(idval, AddComment) {

		for _, arrval := range AddComment {
			for _, j := range arrval {
				if j.Id == idval {
					fetchComment = append(fetchComment, Comment{
						Id:       j.Id,
						Postid:   j.Postid,
						CommPost: j.CommPost,
					})
				}
			}

		}
		w.Header().Set("Content-Type", "aplication/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(fetchComment)

	} else {
		HandleNotFoundError(w, r)
	}

}

func getPostById(w http.ResponseWriter, r *http.Request, id string) {
	idval, err := strconv.Atoi(id)
	if HandleStringConvError(err) {
		return
	}

	val, ok := Addpost[idval]
	if !ok {
		HandleNotFoundError(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(val)

}
func detelePostById(w http.ResponseWriter, r *http.Request, splitstr []string) {
	idval, err := strconv.Atoi(splitstr[2])
	if HandleStringConvError(err) {
		return
	}
	if HandlePostPresent(idval, Addpost) {
		delete(Addpost, idval)
		DeleteCommentByPostId(idval)
	} else {
		HandleNotFoundError(w, r)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(Addpost)

}
func DeleteCommentByPostId(idval int) {
	for i, j := range AddComment {
		if j[i].Id == idval {
			delete(AddComment, j[i].Id)
		}
	}
}
func UpdatePostById(w http.ResponseWriter, r *http.Request, splitstring []string) {
	idval, err := strconv.Atoi(splitstring[2])
	if HandleStringConvError(err) {
		return
	}
	_, ok := Addpost[idval]
	if !ok {
		HandleNotFoundError(w, r)
		return
	}
	var post postData
	json.NewDecoder(r.Body).Decode(&post)
	updatedContent := Addpost[idval]
	updatedContent.Caption = post.Caption
	Addpost[idval] = updatedContent
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	json.NewEncoder(w).Encode(Addpost)

}
func handlePostFunc(w http.ResponseWriter, r *http.Request) {

	stringval := strings.TrimPrefix(r.URL.Path, "/posts")
	splitstring := strings.Split(stringval, "/")
	switch r.Method {

	case http.MethodGet:
		// get post/{id}/comment
		if len(splitstring) == 4 && splitstring[3] == "comment" {
			fetchCommentById(w, r, splitstring)
		}
		// get /post/{id}
		if len(splitstring) == 3 {
			getPostById(w, r, splitstring[2])
		}

	case http.MethodPost:
		//post /post/{id}/comment
		if len(splitstring) == 4 && splitstring[3] == "comment" {
			addCommentById(w, r, splitstring)
		}
		//post /post/{id}
		if len(splitstring) == 3 {
			addCommentById(w, r, splitstring)
		}
		if len(splitstring) == 2 {
			createPost(w, r)
		}
	case http.MethodPut:
		//put /post/{id}
		if len(splitstring) == 3 {
			UpdatePostById(w, r, splitstring)
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
