package webapp

import (
	"fmt"
	"go-course-4/homework-12/server/pkg/index"
	"net/http"
)

const addr = "0.0.0.0:8080"

// Handler -
func Handler(index *index.Index) {
	mux := http.NewServeMux()

	mux.HandleFunc("/docs", handlerDocs)
	mux.HandleFunc("/index", handlerIndex)

	err := http.ListenAndServe(addr, mux)
	if err != nil {
		fmt.Println("Error", err)
	}
}

func handlerDocs(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/docs" {
		handler404(w)
		return
	}
	w.Write([]byte("Docs"))
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/index" {
		handler404(w)
		return
	}
	w.Write([]byte("Index"))
}

func handler404(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Page Not Found"))
}
