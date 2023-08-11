package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"go-course-4/homework-13/pkg/crawler"
	"go-course-4/homework-13/pkg/index"
)

// API - служба API.
type API struct {
	router *mux.Router
	index  *index.Index
}

// New - конструктор для API.
func New(index *index.Index, mux *mux.Router) *API {
	api := API{
		router: mux,
		index:  index,
	}

	api.endpoints()
	api.router.Use(headersMiddleware)
	return &api
}

func (api *API) endpoints() {
	api.router.HandleFunc("/api/v1/docs/{word}", api.searchDoc).Methods(http.MethodGet)
	api.router.HandleFunc("/api/v1/docs", api.createDoc).Methods(http.MethodPost)
	api.router.HandleFunc("/api/v1/docs/{id}", api.updateDoc).Methods(http.MethodPut)
	api.router.HandleFunc("/api/v1/docs/{id}", api.destroyDoc).Methods(http.MethodDelete)
}

func (api *API) searchDoc(w http.ResponseWriter, r *http.Request) {
	data := api.index.Search(mux.Vars(r)["word"])
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (api *API) createDoc(w http.ResponseWriter, r *http.Request) {
	var doc crawler.Document
	err := json.NewDecoder(r.Body).Decode(&doc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var docs []crawler.Document
	docs = append(docs, doc)
	api.index.AddDocuments(docs)
}

func (api *API) updateDoc(w http.ResponseWriter, r *http.Request) {
	//id, err := strconv.Atoi(mux.Vars(r)["id"])
}

func (api *API) destroyDoc(w http.ResponseWriter, r *http.Request) {
	//id, err := strconv.Atoi(mux.Vars(r)["id"])
}
