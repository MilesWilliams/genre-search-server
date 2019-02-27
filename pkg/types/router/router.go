package router

import (
	"net/http"
)

// Routes slice
type Routes []Route

// Route struct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// SubRoutePackage struct
type SubRoutePackage struct {
	Routes     Routes
	Middleware func(next http.Handler) http.Handler
}
