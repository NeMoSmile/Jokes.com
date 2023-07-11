package home

import (
	"net/http"
	"text/template"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("username")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	tmpl := template.Must(template.ParseFiles("view/view/main.html"))
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
