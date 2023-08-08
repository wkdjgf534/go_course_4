package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"go-course-4/homework-13/pkg/crawler"
	"go-course-4/homework-13/pkg/index"
)

// API -
type API struct {
	router *mux.Router
	index  *index.Index
}

// New -
func New(index *index.Index) *API {
	api := API{
		router: mux.NewRouter(),
		index:  index,
	}

	api.endpoints()
	return &api
}

func (api *API) endpoints() {
	// Middleware Header w.Header().Set("Content-Type", "application/json")

	api.router.HandleFunc("/api/v1/docs", api.docs).Methods(http.MethodGet)
	api.router.HandleFunc("/api/v1/docs/{id}", api.showDoc).Methods(http.MethodGet)
	api.router.HandleFunc("/api/v1/docs", api.createDoc).Methods(http.MethodPost)
	api.router.HandleFunc("/api/v1/docs/{id}", api.updateDoc).Methods(http.MethodPut)
	api.router.HandleFunc("/api/v1/docs/{id}", api.destroyDoc).Methods(http.MethodDelete)
}

func (api *API) docs(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(index.docs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (api *API) showDocs(w http.ResponseWriter, r *http.Request) {
	//id, err := strconv.Atoi(mux.Vars(r)["id"])
}

func (api *API) createDoc(w http.ResponseWriter, r *http.Request) {
	var doc crawler.Document
	err := json.NewDecoder(r.Body).Decode(&doc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//api.index.AddDocuments(doc)
}

func (api *API) updateDoc(w http.ResponseWriter, r *http.Request) {
	//id, err := strconv.Atoi(mux.Vars(r)["id"])
}

func (api *API) destroyDoc(w http.ResponseWriter, r *http.Request) {
	//id, err := strconv.Atoi(mux.Vars(r)["id"])
}
