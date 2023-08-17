package webapp

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"go-course-4/homework-12/pkg/crawler"
	"go-course-4/homework-12/pkg/index"
)

var testMux *mux.Router

func TestMain(m *testing.M) {
	docs := []crawler.Document{{ID: 0, Title: "Document0"}}
	testMux = mux.NewRouter()
	index := index.New()
	index.AddDocuments(docs)

	wa := New(index)
	testMux.HandleFunc("/docs", wa.DocsHandler).Methods(http.MethodGet)
	testMux.HandleFunc("/index", wa.IndexHandler).Methods(http.MethodGet)
	m.Run()
}

func TestDocsHandler(t *testing.T) {
	want := "<html><body><div><p>0 - Document0</p></div></body></html>"

	req := httptest.NewRequest(http.MethodGet, "/docs", nil)
	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "HTTP status is 200")

	data, err := io.ReadAll(rr.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	got := string(data)

	assert.Equal(t, want, got, "got expected HTML template")
}

func TestIndexHandler(t *testing.T) {
	want := "<html><body><div><p>0 - Document0</p></div></body></html>"

	req := httptest.NewRequest(http.MethodGet, "/docs", nil)

	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "HTTP status is OK")

	data, err := io.ReadAll(rr.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	got := string(data)

	assert.Equal(t, want, got, "got expected HTML template")
}
