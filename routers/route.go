package routers

import "net/http"

// Route sdfd s
type Route struct {
	Name    string
	Pattern string
	Method  string
	Handler http.HandlerFunc
}

// Routes is a collections of route
type Routes []Route
