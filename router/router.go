package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/my/go-sample/article"
)

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := article.Articles{}
	for i := 0; i < 5; i++ {
		title := "Hello_%d"
		articles = append(
			articles,
			article.Article{Title: fmt.Sprintf(title, i), Desc: "Article Description", Content: "Article Content"})
	}
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage!")
	fmt.Println("Endpoint Hit: homePage")
}

// HandleRequests です
func HandleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
