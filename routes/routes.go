package routes

import (
	"http"
	"github.com/gorilla/mux"
	c "github.com/skantuz/backend/controllers"
)

type Route struct {
	Name      string
	Method    string
	Patther   string
	HadleFunc http.HadlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	fs = http.FileServer(http.Dir("./public"))
	for _, route := range routes{
		r.
		Methods(route.Method).
		Path(route.Patther).
		Handler(route.HandleFunc)
	}
	return r
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/api",
		c.Index,
	}
}
