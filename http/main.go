package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func noteList(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/templates/home.html")
	if err != nil {
		http.Error(w, "An error had been found", http.StatusInternalServerError)
	}
	t.Execute(w, nil)
}
func noteView(w http.ResponseWriter, r *http.Request) {

}
func noteCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		//reject the request
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w, "Creating notes")

}

func main() {
	fmt.Println("Rodando porta 5000")
	mux := http.NewServeMux()

	mux.HandleFunc("/", noteList)
	mux.HandleFunc("/note/view", noteView)
	mux.HandleFunc("/note/create", noteCreate)
	http.ListenAndServe(":5000", mux)

}
