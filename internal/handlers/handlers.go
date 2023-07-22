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

	if len(name) > 12 || strings.Contains(password, " ") || len(password) < 5 || len(password) > 110 || len(email) > 100 {
		http.Redirect(w, r, "/registr", http.StatusFound)
		return
	}

	if d.Check(email, password) == 3 {
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
	ch := d.Check(email, password)
	if ch == 1 {
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
	if ch == 3 {
		http.Redirect(w, r, "/registr", http.StatusFound)
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
	cookie, err := r.Cookie("username")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	email := cookie.Value

	var allW []string = d.WData(email)

	pageContent := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>W</title>
		<style>
			body {
				text-align: center;
				font-size: 30px;
				width: 100%;
				display: flex;
				justify-content: center;
			}
			
			.content {
				position: fixed;
				top: 1%;
				right: 15%;
				height: 97%;
				width: 70%;
				border: 1px solid #ddd;
				border-radius: 10px;
				overflow: hidden;
				box-shadow: 0 2px 10px 0 rgba(255, 214, 247, 0.868);
			}
			
			.content ul {
				list-style-type: none;
				padding: 20px;
				margin: 0;
				overflow-y: scroll;
				height: 100%; 
			}
			
			.content li {
				background-color: #e4ffe3;
				border-radius: 10px;
				padding: 20px;
				margin-bottom: 20px;
				font-weight: 300; 
			}
		</style>
		<script>document.addEventListener('DOMContentLoaded', function () {
			const url = window.location.href;
			window.history.replaceState(null, null, url);
		});</script>
	</head>
	<body>
		<div class="content">
			<ul>
				<li>Hello, how are you?</li>`
	for _, element := range allW {
		pageContent += "<li>" + element + "</li>"
	}
	pageContent += `
	</ul>
</div>

</body>
</html>
`
	tmpl := template.New("w.html")
	tmpl, err = tmpl.Parse(pageContent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
