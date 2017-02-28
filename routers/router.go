package routers

import (
	"github.com/gorilla/mux"
)

// InitRoutes build Routes dynamically from AllRoutes Array
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	router = SetRoutes(publicRoutes, router)

	routerVersion := router.PathPrefix("/v1").Subrouter()
	routerVersion = SetRoutes(albumsRoutes, routerVersion)
	routerVersion = SetRoutes(artistsRoutes, routerVersion)

	return router
}

// SetRoutes routes to main router
func SetRoutes(routes Routes, router *mux.Router) *mux.Router {
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Handler(route.Handler).
			Name(route.Name)
	}
	return router
}
