package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"go-course-4/homework-13/pkg/crawler"
	"go-course-4/homework-13/pkg/index"
)

// API - служба API.
type API struct {
	index *index.Index
}

// New - конструктор для API.
func New(ind *index.Index) *API {
	return &API{index: ind}
}

// SearchDoc - функция поиска документа по тегу
func (api *API) SearchDoc(w http.ResponseWriter, r *http.Request) {
	data := api.index.Search(mux.Vars(r)["word"])
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// CreateDoc - функция создания лдокумента
func (api *API) CreateDoc(w http.ResponseWriter, r *http.Request) {
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

// UpdateDoc - функция обновления документам по id документа
func (api *API) UpdateDoc(w http.ResponseWriter, r *http.Request) {
	var doc crawler.Document
	err := json.NewDecoder(r.Body).Decode(&doc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	status := api.index.ExistsDoc(id)
	if status {
		api.index.Docs[id].URL = doc.URL
		api.index.Docs[id].Title = doc.Title
		api.index.Docs[id].Body = doc.Body
	} else {
		http.Error(w, "Document was not found", http.StatusNotFound)
	}
}

// DestroyDoc - функция удаления документа по id
func (api *API) DestroyDoc(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	status := api.index.ExistsDoc(id)
	if status {
		api.index.Docs = append(api.index.Docs[:id], api.index.Docs[id+1:]...)
	} else {
		http.Error(w, "Document was not found", http.StatusNotFound)
	}

}
