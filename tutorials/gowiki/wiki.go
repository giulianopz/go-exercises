package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type Page struct {
	Title string
	Body  []byte
}

type Home struct {
	Title   string
	Entries []string
}

var templates = template.Must(template.ParseFiles("./tmpl/edit.html", "./tmpl/view.html", "./tmpl/home.html"))

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func (p *Page) save() error {
	filename := "data/" + strings.ReplaceAll(p.Title, " ", "") + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := "data/" + title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func newHandler(w http.ResponseWriter, r *http.Request) {

	title := r.FormValue("ftitle")
	if title == "" {
		http.Error(w, "missing title", http.StatusInternalServerError)
		return
	}

	files, readErr := ioutil.ReadDir("./data/")
	if readErr != nil {
		log.Fatal(readErr)
	}

	suffix := ".txt"
	for _, file := range files {
		if strings.HasSuffix(file.Name(), suffix) && title == strings.TrimSuffix(file.Name(), ".txt") {
			editHandler(w, r, title)
		}
	}

	p := &Page{Title: title}
	renderTemplate(w, "edit", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func home(w http.ResponseWriter, r *http.Request) {

	files, readErr := ioutil.ReadDir("./data/")
	if readErr != nil {
		log.Fatal(readErr)
	}

	suffix := ".txt"
	entries := make([]string, 0)
	for _, file := range files {
		if strings.HasSuffix(file.Name(), suffix) {
			entries = append(entries, strings.TrimSuffix(file.Name(), ".txt"))
		}
	}

	p := &Home{Title: "home", Entries: entries}
	templErr := templates.ExecuteTemplate(w, "home.html", p)
	if templErr != nil {
		http.Error(w, templErr.Error(), http.StatusInternalServerError)
	}
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/new/", newHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
