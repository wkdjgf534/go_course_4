package webapp

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	"go-course-4/homework-12/pkg/crawler"
	"go-course-4/homework-12/pkg/index"
)

var testMux *mux.Router

func TestMain(m *testing.M) {
	docs := []crawler.Document{{ID: 0, Title: "Document0"}}
	index := *index.New()
	index.AddDocuments(docs)
	data = &index

	testMux = mux.NewRouter()
	endpoints(testMux)
	m.Run()
}

func Test_docsHandler(t *testing.T) {
	want := "<html><body><div><p>0 - Document0</p></div></body></html>"

	req := httptest.NewRequest(http.MethodGet, "/docs", nil)

	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("got HTTP code: %d, want HTTP code: %d", rr.Code, http.StatusOK)
	}

	data, err := io.ReadAll(rr.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	got := string(data)

	if got != want {
		t.Errorf("got HTML code %v, want HTML code %v", got, want)
	}
}

func Test_indexHandler(t *testing.T) {
	want := "<html><body><div><p>0 - Document0</p></div></body></html>"

	req := httptest.NewRequest(http.MethodGet, "/docs", nil)

	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("got HTTP code: %d, want HTTP code: %d", rr.Code, http.StatusOK)
	}

	data, err := io.ReadAll(rr.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	got := string(data)

	if got != want {
		t.Errorf("got HTML code %v, want HTML code %v", got, want)
	}
}
