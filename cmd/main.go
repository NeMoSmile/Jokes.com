package main

import (
	"html/template"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/", hrHandler)
	mux.HandleFunc("/registr", registrHandler)
	mux.HandleFunc("/login", hrHandler)
	mux.HandleFunc("/auth", hrHandler)
	mux.HandleFunc("/what", whatHandler)
	mux.HandleFunc("/w", wHandler)
	http.ListenAndServe(":8080", mux)
}

func hrHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("view/authentication/login.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func registrHandler(w http.ResponseWriter, r *http.Request) {
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
