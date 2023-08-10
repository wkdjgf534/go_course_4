package webapp

import (
	"fmt"
	"net/http"

	"go-course-4/homework-13/pkg/api"
	"go-course-4/homework-13/pkg/index"
)

type WebApp struct {
	index *index.Index
	api   *api.API
}

// New - конструктор для цуи сервера
func New(ind *index.Index) WebApp {
	return WebApp{index: ind}
}

// DocsHandler - функция возвращает список документов
func (webapp *WebApp) DocsHandler(w http.ResponseWriter, r *http.Request) {
	var pTags string
	pTags = "Docs do not exist"

	if len(webapp.index.Docs) != 0 {
		pTags = ""

		for _, doc := range webapp.index.Docs {
			pTags += fmt.Sprintf("<p>%v - %v</p>", doc.ID, doc.Title)
		}
	}
	fmt.Fprintf(w, "<html><body><div>%v</div></body></html>", pTags)
}

// IndexHandler - функция возвращает список тегов и массив id документов
func (webapp *WebApp) IndexHandler(w http.ResponseWriter, r *http.Request) {
	var pTags string
	pTags = "Indexes do not exist"

	if len(webapp.index.Docs) != 0 {
		pTags = ""
		for key, value := range webapp.index.Words {
			pTags += fmt.Sprintf("<p>%v - %v</p>", key, value)
		}
	}
	fmt.Fprintf(w, "<html><body><div>%v</div></body></html>", pTags)
}
