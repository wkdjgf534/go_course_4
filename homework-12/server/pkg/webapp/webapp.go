package webapp

import (
	"fmt"
	"go-course-4/homework-12/server/pkg/index"
	"net/http"

	"github.com/gorilla/mux"
)

const addr = ":8080"

var data *index.Index

// Listen -
func Listen(i *index.Index) {
	data = i

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
	//r.HandleFunc("/index/{id}", indexHandler).Methods(http.MethodGet)
}

func docsHandler(w http.ResponseWriter, r *http.Request) {
	var pTags string
	for _, doc := range data.Docs {
		pTags += fmt.Sprintf("<p>%v - %v</p>", doc.ID, doc.Title)
	}
	fmt.Fprintf(w, "<html><body><div>%v</div></body></html>", pTags)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var pTags string
	for key, value := range data.Words {
		pTags += fmt.Sprintf("<p>%v - %v</p>", key, value)
	}
	fmt.Fprintf(w, "<html><body><div>%v</div></body></html>", pTags)
}
