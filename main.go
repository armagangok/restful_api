package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title       string `json:Title`
	Description string `json:Description`
	Content     string `json:Content`
}

type MyArticle []Article

func allArticle(w http.ResponseWriter, _ *http.Request) {
	articles := Article{
		Title:       "Test Title",
		Description: "Test Description",
		Content:     "Hello, this is an example API for my first backend project.",
	}

	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")

}

func handleRequests() {
	myRouter := mux.newRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticle)
	log.Fatal(http.ListenAndServe(":8081", myRouter))

}

func main() {
	handleRequests()
}
