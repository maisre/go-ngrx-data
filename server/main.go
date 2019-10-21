package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

	r.HandleFunc("/api/post/{postID}", getSinglePost).Methods("GET")
	r.HandleFunc("/api/post/{postID}", updatePost).Methods("PUT")
	r.HandleFunc("/api/post/{postID}", deletePost).Methods("DELETE")
	r.HandleFunc("/api/post/", createPost).Methods("POST")
	r.HandleFunc("/api/post/", getAllPosts).Methods("GET")
	r.HandleFunc("/api/post/", optionReply).Methods("OPTIONS")
	r.HandleFunc("/api/post/{postID}", optionReply).Methods("OPTIONS")

	http.Handle("/", r)

	// Solves Cross Origin Access Issue
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})
	handler := c.Handler(r)

	srv := &http.Server{
		Handler: handler,
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

	t.ID = strconv.Itoa(len(postMap) + 1)
	_, prs := postMap[t.ID]
	if prs {
		fmt.Println("exiting because its there")
		w.WriteHeader(400)
		return
	}

	postMap[t.ID] = t

	w.WriteHeader(200)
	return
}

func getAllPosts(w http.ResponseWriter, r *http.Request) {
	v := make([]post, 0, len(postMap))

	for _, value := range postMap {
		v = append(v, value)
	}

	jsonBytes, err := structToJSON(v)
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

func optionReply(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	return
}
