package main

import (
	"github.com/gorilla/mux"
	"github.com/realisijola/mysql-form/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.Router(r)
	log.Println("starting server at http://localhost:419")
	log.Fatal(http.ListenAndServe(":419", r))
}