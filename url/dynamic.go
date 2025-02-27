package main

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"

	// "os"
	// "path/filepath"
	// "time"

	"github.com/gorilla/mux"
)

func ArticleHandler(w http.ResponseWriter,r *http.Request){
	fmt.Println("hello world")
}


func main(){
	r := mux.NewRouter()

	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler).
    	Name("article") // Assigns the name "article" to this route

	url, err := r.Get("article").URL("category", "tech1", "id", "123")
	if err != nil {
    	log.Fatal(err)
	}

	fmt.Println(url.String()) 
	 // Output: /articles/tech/123
	http.ListenAndServe(":8080",r)

}