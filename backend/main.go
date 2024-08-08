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
	// Serve static files
	jsDir := filepath.Join(frontendDir, "js")
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir(jsDir))))

	routes.RegisterRoutes(r)

	http.ListenAndServe(":8080", r)
}
