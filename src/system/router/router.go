package router

import (
	"net/http"

	"github.com/MilesWilliams/music/pkg/types/router"
	mux "github.com/gorilla/mux"
)

// Router struct
type Router struct {
	Router *mux.Router
}

// Init func
func (r *Router) Init() {
	r.Router.Use(Middleware)
	baseRoutes := GetRoutes()

	for _, route := range baseRoutes {
		r.Router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	// r.Router.NotFoundHandler = http.HandlerFunc(NotFoundHandler.NotFound)
	r.Router.PathPrefix("/public").Handler(http.StripPrefix("/", http.FileServer(http.Dir(""))))
	r.Router.PathPrefix("/static").Handler(http.StripPrefix("/", http.FileServer(http.Dir(""))))
	r.Router.PathPrefix("/static/templates").Handler(http.StripPrefix("/", http.FileServer(http.Dir(""))))
}

// AttachSubRouterWithMiddleware func
func (r *Router) AttachSubRouterWithMiddleware(path string, subroutes router.Routes, middleware mux.MiddlewareFunc) (SubRouter *mux.Router) {
	SubRouter = r.Router.PathPrefix(path).Subrouter()
	SubRouter.Use(middleware)

	for _, route := range subroutes {
		SubRouter.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return
}

// NewRouter func
func NewRouter() (r Router) {
	r.Router = mux.NewRouter().StrictSlash(true)

	return
}
