package main

import (
	"net/http"

	"github.com/adaken4/peerfund-p2p-lending/backend/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	http.ListenAndServe(":8080", r)
}
