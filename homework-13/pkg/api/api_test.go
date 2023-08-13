package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"go-course-4/homework-13/pkg/crawler"
	"go-course-4/homework-13/pkg/index"
)

var api *API
var testMux *mux.Router

func TestMain(m *testing.M) {
	docs := []crawler.Document{{ID: 0, URL: "link_0", Title: "Doc0"}, {ID: 1, URL: "link_1", Title: "Doc1"}}
	testMux = mux.NewRouter()
	index := index.New()
	index.AddDocuments(docs)

	api = New(index)
	testMux.HandleFunc("/api/v1/docs/{word}", api.SearchDoc).Methods(http.MethodGet)
	testMux.HandleFunc("/api/v1/docs", api.CreateDoc).Methods(http.MethodPost)
	testMux.HandleFunc("/api/v1/docs/{id}", api.UpdateDoc).Methods(http.MethodPut)
	testMux.HandleFunc("/api/v1/docs/{id}", api.DestroyDoc).Methods(http.MethodDelete)
	m.Run()
}

func TestSearchDoc(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/docs/go", nil)
	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "HTTP status is 200")
}

func TestCreateDoc(t *testing.T) {
	doc := crawler.Document{URL: "http://gov.uk", Title: "Gov UK"}
	payload, _ := json.Marshal(doc)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/docs", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "HTTP status is 200")

	id := len(api.index.Docs)
	got := api.index.Docs[id-1]
	assert.Equal(t, doc, got, "Created a new correct document")
}

func TestUpdateDoc(t *testing.T) {
	docID := 1
	doc := crawler.Document{ID: docID, URL: "http://gov.uk", Title: "Gov UK"}
	payload, _ := json.Marshal(doc)
	fmt.Println(bytes.NewBuffer(payload))

	req := httptest.NewRequest(http.MethodPut, "/api/v1/docs/"+strconv.Itoa(docID), bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "HTTP status is 200")

	got := api.index.Docs[docID]
	assert.Equal(t, doc, got, "Updated the correct document")
}

func TestDestroyDoc(t *testing.T) {
	docID := 1
	docsAmount := len(api.index.Docs)
	req := httptest.NewRequest(http.MethodDelete, "/api/v1/docs/"+strconv.Itoa(docID), nil)
	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "HTTP status is 200")

	got := len(api.index.Docs)
	want := docsAmount - 1
	assert.Equal(t, want, got, "Deleted the correct document")
}
