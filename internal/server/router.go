package server

import (
	"net/http"
	"path/filepath"

	"github.com/VitorAngelozi/quickNotes/internal/handlers"
)

func NewMux() *http.ServeMux {
	mux := http.NewServeMux()

	staticFiles := http.FileServer(http.Dir(filepath.Join(handlers.ResolveViewsDir(), "static")))
	mux.Handle("/static/", http.StripPrefix("/static/", staticFiles))

	mux.HandleFunc("/", handlers.NoteList)
	mux.HandleFunc("/note/view", handlers.NoteView)
	mux.HandleFunc("/note/create", handlers.NoteCreate)

	return mux
}
