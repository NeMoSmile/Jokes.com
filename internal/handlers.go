package handlers

import (
	"net/http"
	"text/template"
)

func hrHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("view/authentication/registration.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func registrHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	if login == "admin" && password == "12345" {
		// Если логин и пароль верны, ставим куки с именем пользователя
		http.SetCookie(w, &http.Cookie{
			Name:  "username",
			Value: login,
			Path:  "/",
		})

		// Перенаправляем пользователя на защищенную страницу
		http.Redirect(w, r, "/user/", http.StatusFound)
		return
	}

	// Если логин и пароль не верны, возвращаем пользователя на страницу авторизации
	http.Redirect(w, r, "/login", http.StatusFound)
}

func whatHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("view/view/what.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func wHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("view/view/w.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
