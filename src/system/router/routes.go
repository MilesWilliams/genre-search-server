package router

import (
	"net/http"

	"github.com/MilesWilliams/personal/pkg/types/routes"

	SearchHandler "github.com/MilesWilliams/music/src/controllers/search"
)

// Middleware func
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

// GetRoutes func
func GetRoutes() routes.Routes {
	return routes.Routes{
		routes.Route{"Search", "POST", "/search", SearchHandler.Search},
	}
}
