package home

import (
	"fmt"
	"net/http"
	"text/template"

	d "github.com/NeMoSmile/Jokes.com.git/internal/DataBase"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("id")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	id := cookie.Value
	if !d.CheckUser(id) {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	tmpl := template.Must(template.ParseFiles("view/view/main.html"))
	err = tmpl.Execute(w, d.PageData(id))
	if err != nil {
		fmt.Println(err)
	}
}
