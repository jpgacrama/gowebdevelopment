package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"io"
	"log"
	"net/http"
)

const (
	CONN_HOST      = "localhost"
	CONN_PORT      = "8080"
	ADMIN_USER     = "admin"
	ADMIN_PASSWORD = "password"
	REALM          = "Restricted Area"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!")
}

func setupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorld)
	return mux
}

func basicAuthHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()

		if !ok || user != ADMIN_USER || pass != ADMIN_PASSWORD {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+REALM+`"`)
			w.WriteHeader(401)
			w.Write([]byte("Unauthorized. Please provide valid credentials.\n"))
			return
		}

		handler(w, r)
	}
}

func startServer() {
	mux := setupRoutes()
	authHandler := basicAuthHandler(mux.ServeHTTP)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, handlers.CompressHandler(authHandler))
	if err != nil {
		log.Fatal("Error starting http server: ", err)
	}
}

func main() {
	fmt.Printf("Open http://%s:%s/ in your browser.\nClick CTRL + C to exit", CONN_HOST, CONN_PORT)
	startServer()
}
