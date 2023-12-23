package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!")
}

func setupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorld)
	return mux
}

func startServer() {
	mux := setupRoutes()
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, handlers.CompressHandler(mux))
	if err != nil {
		log.Fatal("error starting http server: ", err)
	}
}

func main() {
	fmt.Printf("Open http://%s:%s/ in your browser.\nClick CTRL + C to exit", CONN_HOST, CONN_PORT)
	startServer()
}
