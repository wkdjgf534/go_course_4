package webapp

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"go-course-4/homework-13/pkg/index"
)

const addr = ":8080"

var data *index.Index

// Start - запуск сетевой службы
func Start(i *index.Index) error {
	data = i
	r := mux.NewRouter()
	endpoints(r)
	err := http.ListenAndServe(addr, r)
	if err != nil {
		return err
	}
	return nil
}

func endpoints(r *mux.Router) {
	r.HandleFunc("/docs", docsHandler).Methods(http.MethodGet)
	r.HandleFunc("/index", indexHandler).Methods(http.MethodGet)
}

func docsHandler(w http.ResponseWriter, r *http.Request) {
	var pTags string
	pTags = "Docs do not exist"

	if len(data.Docs) != 0 {
		pTags = ""

		for _, doc := range data.Docs {
			pTags += fmt.Sprintf("<p>%v - %v</p>", doc.ID, doc.Title)
		}
	}
	fmt.Fprintf(w, "<html><body><div>%v</div></body></html>", pTags)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var pTags string
	pTags = "Indexes do not exist"

	if len(data.Docs) != 0 {
		pTags = ""
		for key, value := range data.Words {
			pTags += fmt.Sprintf("<p>%v - %v</p>", key, value)
		}
	}
	fmt.Fprintf(w, "<html><body><div>%v</div></body></html>", pTags)
}
