package routes

import (
	"html/template"
	"net/http"

	"github.com/adaken4/peerfund-p2p-lending/backend/controllers"
	"github.com/gorilla/mux"
)

var templates *template.Template

func init() {
	templates = template.Must(template.ParseGlob("../frontend/templates/*html"))
}

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/login", serveTemplate("login.html")).Methods("GET")
	r.HandleFunc("/register", serveTemplate("register.html")).Methods("GET")

	r.HandleFunc("/register", controllers.RegisterUser).Methods("POST")
	r.HandleFunc("/login", controllers.LoginUser).Methods("POST")
}

func serveTemplate(tmpl string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, tmpl, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
