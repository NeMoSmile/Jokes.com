package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
	"text/template"
	"time"

	d "github.com/NeMoSmile/Jokes.com.git/internal/DataBase"
)

type User struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func StartLoginHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("id")
	if err == nil {
		if d.CheckUser(cookie.Value) {
			http.Redirect(w, r, "/main", http.StatusFound)
			return
		}
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

	if len(name) > 15 || strings.Contains(password, " ") || len(password) > 110 || len(email) > 100 {
		http.Redirect(w, r, "/errorregistr", http.StatusFound)
		return
	}

	if d.Check(email, password) == 3 {

		user := User{
			Email:    email,
			Username: name,
			Password: password,
		}

		userJSON, err := json.Marshal(user)
		if err != nil {
			log.Fatal(err)
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "user",
			Value:    url.QueryEscape(string(userJSON)),
			Expires:  time.Now().Add(1 * time.Hour),
			Path:     "/",
			HttpOnly: true,
		})

		http.Redirect(w, r, "/conf", http.StatusFound)
		return

		// d.Append(email, password, name)

		// http.Redirect(w, r, "/main", http.StatusFound)
		// return
	}
	http.Redirect(w, r, "/", http.StatusFound)

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	ch := d.Check(email, password)
	if ch == 1 {
		id := d.GetId(email)
		http.SetCookie(w, &http.Cookie{
			Name:     "id",
			Value:    id,
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
	if ch == 2 {
		http.Redirect(w, r, "/errorlogin", http.StatusFound)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func WhatHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("id")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	if !d.CheckUser(cookie.Value) {
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

	var allW []string = d.WData(id)

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
				`
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

func ErrorLoginHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("id")
	if err == nil {
		if d.CheckUser(cookie.Value) {
			http.Redirect(w, r, "/main", http.StatusFound)
			return
		}
	}
	tmpl := template.Must(template.ParseFiles("view/authentication/errlogin.html"))
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ErrorRegistrHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("view/authentication/erregistration.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func StartConfirmHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("user")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	encodedUserJSON := cookie.Value
	decodedUserJSON, err := url.QueryUnescape(encodedUserJSON)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	var user User
	err = json.Unmarshal([]byte(decodedUserJSON), &user)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	d.Send(user.Email)
	tmpl := template.Must(template.ParseFiles("view/authentication/confirm.html"))
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ConfirmHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("user")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	encodedUserJSON := cookie.Value
	decodedUserJSON, err := url.QueryUnescape(encodedUserJSON)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	var user User
	err = json.Unmarshal([]byte(decodedUserJSON), &user)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	code := r.FormValue("code")

	if d.CheckUserCode(user.Email, code) {
		id := d.Append(user.Email, user.Password, user.Username)
		http.SetCookie(w, &http.Cookie{
			Name:     "id",
			Value:    id,
			Expires:  time.Now().Add(168 * time.Hour),
			Path:     "/",
			HttpOnly: true,
		})

		http.SetCookie(w, &http.Cookie{
			Name:     "user",
			Value:    "",
			Expires:  time.Now().Add(-1 * time.Hour),
			Path:     "/",
			HttpOnly: true,
		})

		http.Redirect(w, r, "/main", http.StatusFound)
		return
	}
	http.Redirect(w, r, "/conf", http.StatusFound)

}
