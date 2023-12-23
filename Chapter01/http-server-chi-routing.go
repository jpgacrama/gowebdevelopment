package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(fmt.Sprintf("Route %s %s not found", r.Method, r.URL.Path)))
}

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte(fmt.Sprintf("Route %s %s is not allowed", r.Method, r.URL.Path)))
}

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", GetRequestHandler)
	router.Post("/post", PostRequestHandler)
	router.Put("/hello/{name}", PathVariableHandler)

	// Default route handler for unrecognized routes
	router.NotFound(NotFoundHandler)
	router.MethodNotAllowed(MethodNotAllowedHandler)

	fmt.Printf("Server listening on %s:%s\n", CONN_HOST, CONN_PORT)
	http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
}
