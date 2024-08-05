package main

import (
	"net/http"
	"path/filepath"

	"github.com/adaken4/peerfund-p2p-lending/backend/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Serve static files
	frontendDir := filepath.Join("..", "frontend")
	cssDir := filepath.Join(frontendDir, "css")
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir(cssDir))))

	routes.RegisterRoutes(r)

	http.ListenAndServe(":8080", r)
}
