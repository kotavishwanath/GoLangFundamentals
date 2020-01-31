package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Title", Desc: "Test Description", Content: "Welcome to the world of Goooooooooooo"},
		Article{Title: "Golang", Desc: "I am going to be building a fully-fledged REST API", Content: "In order to keep this simple and focus on the basic concepts, we won’t be interacting with any backend database technologies to store the articles that we’ll be playing with. However, we will be writing this REST API in such a way that it will be easy to update the functions we will be defining so that they make subsequent calls to a database to perform any necessary CRUD operations."},
	}
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HomePage Endpoint Hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8083", nil))
}

func main() {
	handleRequests()
}
