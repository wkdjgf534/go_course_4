package webapp

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"go-course-4/homework-12/server/pkg/index"
)

const addr = ":8080"

// Handler -
func Handler(index *index.Index) {
	r := mux.NewRouter()
	endpoints(r)
	err := http.ListenAndServe(addr, r)
	if err != nil {
		fmt.Println("Error", err)
	}
}

func endpoints(r *mux.Router) {
	r.HandleFunc("/docs", docsHandler).Methods(http.MethodGet)
	r.HandleFunc("/index", indexHandler).Methods(http.MethodGet)
}

func docsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Docs"))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index"))
}
