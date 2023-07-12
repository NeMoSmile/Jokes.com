package handlers

import (
	"net/http"
	"strings"
	"text/template"
	"time"

	d "github.com/NeMoSmile/Jokes.com.git/internal/DataBase"
)

func StartLoginHandler(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("username")
	if err == nil {
		http.Redirect(w, r, "/main", http.StatusFound)
		return
	}
	tmpl := template.Must(template.ParseFiles("view/authentication/login.html"))
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func StartRegistrHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("view/authentication/registration.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RegistrHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	if len(name) > 12 || strings.Contains(password, " ") {
		http.Redirect(w, r, "/registr", http.StatusFound)
		return
	}

	if d.Check(email, password) == 1 {
		// Если логин и пароль верны, ставим куки с именем пользователя
		http.SetCookie(w, &http.Cookie{
			Name:     "username",
			Value:    email,
			Expires:  time.Now().Add(168 * time.Hour),
			Path:     "/",
			HttpOnly: true,
		})

		d.Append(email, password, name)

		// Перенаправляем пользователя на защищенную страницу
		http.Redirect(w, r, "/main", http.StatusFound)
		return
	}

	// Если логин и пароль не верны, возвращаем пользователя на страницу авторизации
	http.Redirect(w, r, "/registr", http.StatusFound)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	if d.Check(email, password) == 1 {
		http.SetCookie(w, &http.Cookie{
			Name:     "username",
			Value:    email,
			Expires:  time.Now().Add(168 * time.Hour),
			Path:     "/",
			HttpOnly: true,
		})

		http.Redirect(w, r, "/main", http.StatusFound)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func WhatHandler(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("username")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	tmpl := template.Must(template.ParseFiles("view/view/what.html"))
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func WHandler(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("username")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	tmpl := template.Must(template.ParseFiles("view/view/w.html"))
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
