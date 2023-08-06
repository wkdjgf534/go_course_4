package webapp

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/exp/slices"

	"go-course-4/homework-12/pkg/index"
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
	r.HandleFunc("/index/{word}", indexHandlerByWord).Methods(http.MethodGet)
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

func indexHandlerByWord(w http.ResponseWriter, r *http.Request) {
	indexWord := mux.Vars(r)["word"]
	indexKeys := []string{}

	for k, _ := range data.Words {
		indexKeys = append(indexKeys, k)
	}

	pTags := "Index does not exist"

	if slices.Contains(indexKeys, indexWord) {
		pTags = ""
		for _, doc := range data.Docs {
			pTags += fmt.Sprintf("<p>%v - %v</p>", doc.ID, doc.Title)
		}
	}
	fmt.Fprintf(w, "<html><body><div>%v</div></body></html>", pTags)
}
