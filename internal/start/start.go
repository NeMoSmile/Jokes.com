package start

import (
	"net/http"

	h "github.com/NeMoSmile/Jokes.com.git/internal/handlers"
)

func Start(port string) {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", h.StartLoginHandler)
	mux.HandleFunc("/registr", h.StartRegistrHandler)
	mux.HandleFunc("/login", h.LoginHandler)
	mux.HandleFunc("/auth", h.RegistrHandler)
	mux.HandleFunc("/main", h.MainHandler)
	mux.HandleFunc("/what", h.WhatHandler)
	mux.HandleFunc("/w", h.WHandler)

	http.ListenAndServe(":"+port, mux)
}
