package main

import (
	"html/template"
	"log"
	"net/http"
	"regexp"
)

type Page struct {
	Title string
	Text  string
}

func homeViewHandler(w http.ResponseWriter, r *http.Request, text string) {
	p := &Page{Title: "Home", Text: text}
	renderTemplate(w, "home", p)
}

var templates = template.Must(template.ParseFiles("home.htmltmpl"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".htmltmpl", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var validPath = regexp.MustCompile("^/(home)/([a-zA-Z0-9]+)$")

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

func main() {
	http.HandleFunc("/home/", makeHandler(homeViewHandler))
	log.Fatal(http.ListenAndServe(":80", nil))
}
