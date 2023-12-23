package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

func GetRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func PostRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("It's a Post Request!"))
}

func PathVariableHandler(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	w.Write([]byte("Hi " + name))
}

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", GetRequestHandler)
	router.Post("/post", PostRequestHandler)
	router.Put("/hello/{name}", PathVariableHandler)

	http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
}
