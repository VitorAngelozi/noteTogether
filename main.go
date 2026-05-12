package main

import (
	"fmt"
	"net/http"
)

func noteList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Listing notes")
}
func noteView(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Showing notes")

}
func noteCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow", "POST")

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
