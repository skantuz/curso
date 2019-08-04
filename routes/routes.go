package routes

import (
	"net/http"
	"path"
	"strings"

	"github.com/gorilla/mux"
	c "github.com/skantuz/backend/controllers"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	fs := http.FileServer(http.Dir("./public"))
	for _, route := range routes {
		r.
			Methods(route.Method).
			Path(route.Pattern).
			Handler(route.HandleFunc)
	}
	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		_path := req.URL.Path
		if strings.Contains(_path, ".") || _path == "/" {
			fs.ServeHTTP(w, req)
			return
		}
		http.ServeFile(w, req, path.Join("./public", "/index.html"))
	})
	return r
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/api",
		c.Index,
	},
}
