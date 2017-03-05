package routers

import (
	"github.com/gorilla/mux"
	"net/http"
)

// InitRoutes build Routes dynamically from AllRoutes Array
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	subRouterApiVersion := router.PathPrefix("/v1").Subrouter()

	// Public routes --> without token
	subRouterApiVersion = SetRoutes(publicRoutes, router)

	// Private routes --> with token
	subRouterApiVersion = SetRoutes(albumsRoutes, subRouterApiVersion)
	subRouterApiVersion = SetRoutes(artistsRoutes, subRouterApiVersion)

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

func SetMiddleware(middleware http.Handler, router *mux.Router) *mux.Router {
	router.Handle("", middleware)

}
