// Main package
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"go-course-4/homework-13/pkg/api"
	"go-course-4/homework-13/pkg/crawler"
	"go-course-4/homework-13/pkg/crawler/spider"
	"go-course-4/homework-13/pkg/index"
	"go-course-4/homework-13/pkg/middleware"
	"go-course-4/homework-13/pkg/webapp"
)

const (
	addr  = ":8080"
	depth = 1
)

var urls = []string{"https://golang.org", "https://www.practical-go-lessons.com/"}

func main() {
	s := spider.New()
	ind := index.New()

	var docs []crawler.Document
	for _, u := range urls {
		links, err := s.Scan(u, depth)
		if err != nil {
			fmt.Printf("We got an error: %s\n", err)
			continue
		}
		docs = append(docs, links...)
	}
	ind.AddDocuments(docs)

	mux := mux.NewRouter()
	api := api.New(ind)
	wa := webapp.New(ind)
	apiEndpoints(api, mux)
	webappEndpoints(wa, mux)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		fmt.Printf("Error from http server: %s", err)
		return
	}

}

func webappEndpoints(wa *webapp.WebApp, mux *mux.Router) {
	mux.HandleFunc("/docs", wa.DocsHandler).Methods(http.MethodGet)
	mux.HandleFunc("/index", wa.IndexHandler).Methods(http.MethodGet)
}

func apiEndpoints(api *api.API, mux *mux.Router) {
	mux.Use(middleware.HeadersMiddleware)
	mux.HandleFunc("/api/v1/docs/{word}", api.SearchDoc).Methods(http.MethodGet)
	mux.HandleFunc("/api/v1/docs", api.CreateDoc).Methods(http.MethodPost)
	mux.HandleFunc("/api/v1/docs/{id}", api.UpdateDoc).Methods(http.MethodPut)
	mux.HandleFunc("/api/v1/docs/{id}", api.DestroyDoc).Methods(http.MethodDelete)
}
