package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

type PageData struct {
	Title string
	Note  string
}

func renderTemplate(w http.ResponseWriter, page string, data PageData) {
	baseDir := resolveViewsDir()
	files := []string{
		filepath.Join(baseDir, "templates", "layouts", "base.html"),
		filepath.Join(baseDir, "templates", "partials", "brand.html"),
		filepath.Join(baseDir, "templates", "pages", page),
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "failed to load template", http.StatusInternalServerError)
		return
	}

	if err := t.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, "failed to render template", http.StatusInternalServerError)
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

func ResolveViewsDir() string {
	return resolveViewsDir()
}

func NoteList(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html", PageData{Title: "Note Together"})
}

func NoteView(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "note-view.html", PageData{
		Title: "Note Together | View Note",
		Note:  "",
	})
}

func NoteCreate(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "note-create.html", PageData{Title: "Note Together | Create Note"})
}
