package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var postMap map[string]post

type post struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func newPost(id string) post {
	return post{
		ID:      id,
		Title:   "title",
		Content: "content",
	}
}

func main() {
	postMap = make(map[string]post)
	r := mux.NewRouter()

	postMap["1"] = newPost("1")
	postMap["2"] = newPost("2")

	r.HandleFunc("/post/{postID}", getSinglePost).Methods("GET")
	r.HandleFunc("/post/{postID}", updatePost).Methods("PUT")
	r.HandleFunc("/post/{postID}", deletePost).Methods("DELETE")
	r.HandleFunc("/post/", createPost).Methods("POST")
	r.HandleFunc("/post/", getAllPosts).Methods("GET")

	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    ":" + os.Getenv("PORT"),
	}

	log.Fatal(srv.ListenAndServe())
}

func getSinglePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if len(vars) > 0 {
		p, prs := postMap[vars["postID"]]
		if prs {
			jsonBytes, err := structToJSON(p)
			if err != nil {
				fmt.Print(err)
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonBytes)
			return
		}

		w.WriteHeader(400)
		return

	}

	w.WriteHeader(400)
	return
}
func updatePost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t post
	err := decoder.Decode(&t)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		return
	}

	_, prs := postMap[t.ID]
	if prs {
		fmt.Println("its there")

		postMap[t.ID] = t

		w.WriteHeader(200)
		return
	}

	w.WriteHeader(400)
	return
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if len(vars) > 0 {
		delete(postMap, vars["postID"])
		w.WriteHeader(200)
		return
	}

	w.WriteHeader(400)
	return
}

func createPost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t post
	err := decoder.Decode(&t)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		return
	}

	_, prs := postMap[t.ID]
	if prs {
		w.WriteHeader(400)
		return
	}

	postMap[t.ID] = t

	w.WriteHeader(200)
	return
}

func getAllPosts(w http.ResponseWriter, r *http.Request) {

	jsonBytes, err := structToJSON(postMap)
	if err != nil {
		fmt.Print(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
	return
}

func structToJSON(data interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := json.NewEncoder(buf).Encode(data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
