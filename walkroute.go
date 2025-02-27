package main

// import (
// 	"fmt"
// 	"net/http"
// 	"strings"

// 	"github.com/gorilla/mux"
// )

// func handler(w http.ResponseWriter, r *http.Request) {
// 	return
// }

// func main() {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/", handler)
// 	r.HandleFunc("/products", handler).Methods("POST")
// 	r.HandleFunc("/articles", handler).Methods("GET").Queries("surname","{id}").Schemes("https")
// 	r.HandleFunc("/articles/{id}", handler).Methods("GET", "PUT")
// 	r.HandleFunc("/authors", handler).Queries("surname", "{surname}")
// 	// /author/?surname=aggarwal
// 	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
// 		//fmt.Println(*route)
// 		//fmt.Println(*router)
// 		//fmt.Println(ancestors)
// 		pathTemplate, err := route.GetPathTemplate()
// 		if err == nil {
// 			fmt.Println("ROUTE:", pathTemplate)
// 		}
// 		pathRegexp, err := route.GetPathRegexp()
// 		if err == nil {
// 			fmt.Println("Path regexp:", pathRegexp)
// 		}
// 		queriesTemplates, err := route.GetQueriesTemplates()
// 		if err == nil {
// 			fmt.Println("Queries templates:", strings.Join(queriesTemplates, ","))
// 		}
// 		queriesRegexps, err := route.GetQueriesRegexp()
// 		if err == nil {
// 			fmt.Println("Queries regexps:", strings.Join(queriesRegexps, ","))
// 		}
// 		methods, err := route.GetMethods()
// 		if err == nil {
// 			fmt.Println("Methods:", strings.Join(methods, ","))
// 		}
// 		schema,err := route.Get
// 		fmt.Println()
// 		return nil
// 	})

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	http.Handle("/", r)
// }
//