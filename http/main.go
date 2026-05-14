package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

type pageData struct {
	Title string
	Note  string
}

func renderTemplate(w http.ResponseWriter, page string, data pageData) {
	baseDir := resolveViewsDir()
	files := []string{
		filepath.Join(baseDir, "templates", "layouts", "base.html"),
		filepath.Join(baseDir, "templates", "partials", "brand.html"),
		filepath.Join(baseDir, "templates", "pages", page),
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "An error had been found", http.StatusInternalServerError)
		return
	}

	if err := t.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, "An error had been found", http.StatusInternalServerError)
	}
}

func resolveViewsDir() string {
	candidates := []string{"views", "../views"}
	for _, c := range candidates {
		if info, err := os.Stat(c); err == nil && info.IsDir() {
			return c
		}
	}
	return "views"
}

func noteList(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html", pageData{Title: "Note Together"})
}

func noteView(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "note-view.html", pageData{
		Title: "Note Together | View Note",
		Note:  "",
	})
}

func noteCreate(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "note-create.html", pageData{Title: "Note Together | Create Note"})
}

func main() {
	fmt.Println("Rodando porta 5000")
	mux := http.NewServeMux()

	staticFiles := http.FileServer(http.Dir(filepath.Join(resolveViewsDir(), "static")))
	mux.Handle("/static/", http.StripPrefix("/static/", staticFiles))

	mux.HandleFunc("/", noteList)
	mux.HandleFunc("/note/view", noteView)
	mux.HandleFunc("/note/create", noteCreate)
	http.ListenAndServe(":5000", mux)

}
