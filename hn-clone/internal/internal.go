package internal

import (
	"example/internal/api"
	"example/internal/model"
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseFiles("../../web/index.html"))

func Home(w http.ResponseWriter, r *http.Request) {

	stories, err := api.NextPage(0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	p := &model.Page{Stories: stories}
	templErr := templates.ExecuteTemplate(w, "index.html", p)
	if templErr != nil {
		http.Error(w, templErr.Error(), http.StatusInternalServerError)
	}
}
