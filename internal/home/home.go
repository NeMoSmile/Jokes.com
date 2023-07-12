package home

import (
	"net/http"
	"text/template"

	d "github.com/NeMoSmile/Jokes.com.git/internal/DataBase"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	email := cookie.Value

	tmpl := template.Must(template.ParseFiles("view/view/main.html"))
	err = tmpl.Execute(w, d.PageData(email))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
