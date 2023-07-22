package start

import (
	"net/http"

	h "github.com/NeMoSmile/Jokes.com.git/internal/handlers"
	hm "github.com/NeMoSmile/Jokes.com.git/internal/home"
)

func Start(port string) {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// GET page for user authorization
	mux.HandleFunc("/", h.StartLoginHandler)

	// GET page for registration.
	mux.HandleFunc("/registr", h.StartRegistrHandler)

	// POST page for authorization.
	mux.HandleFunc("/login", h.LoginHandler)

	// POST page for registration.
	mux.HandleFunc("/auth", h.RegistrHandler)

	// Main page
	mux.HandleFunc("/main", hm.MainHandler)

	// Question mark page
	mux.HandleFunc("/what", h.WhatHandler)

	// Page with tagged jokes
	mux.HandleFunc("/w", h.WHandler)

	mux.HandleFunc("/errorlogin", h.ErrorLoginHandler)

	mux.HandleFunc("/errorregistr", h.ErrorRegistrHandler)

	http.ListenAndServe(port, mux)
}
