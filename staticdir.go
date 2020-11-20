package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := NewRouter()
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Choose the folder to serve
	staticDir := "/static/"

	// Create the route
	router.
		PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	return router
}
