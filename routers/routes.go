package routers

import (
	"github.com/christiangda/gogo-rest-api/handlers"
	"net/http"
)

// Route sdfd s
type Route struct {
	Name    string
	Pattern string
	Method  string
	Handler http.HandlerFunc
}

// Routes is a collections of route
type Routes []Route

// Public:  Routes that doesn't need Authorization
var publicRoutes = Routes{
	{Name: "Index", Pattern: "/", Method: http.MethodGet, Handler: handlers.Index},
}

// Public: Users routes
var usersRoutes = Routes{
	{Name: "Register", Pattern: "/users/register", Method: http.MethodGet, Handler: handlers.Register},
	{Name: "Login", Pattern: "/users/login", Method: http.MethodPost, Handler: handlers.Login},
	{Name: "Me", Pattern: "/users/me", Method: http.MethodGet, Handler: handlers.Me},
}

// Privated: Albums Routes
var albumsRoutes = Routes{
	{Name: "Album", Pattern: "/albums/{id}", Method: http.MethodGet, Handler: handlers.Album},
	{Name: "Albums", Pattern: "/albums?ids={ids}", Method: http.MethodGet, Handler: handlers.Albums},
	{Name: "AlbumTracks", Pattern: "/albums/{id}/tracks", Method: http.MethodGet, Handler: handlers.AlbumTracks},
}

// Privated: Artists Routes
var artistsRoutes = Routes{
	{Name: "Artist", Pattern: "/artists/{id}", Method: http.MethodGet, Handler: handlers.Artist},
	{Name: "Artists", Pattern: "/artists?ids={ids}", Method: http.MethodGet, Handler: handlers.Artists},
	{Name: "ArtistAlbums", Pattern: "/artists/{id}/albums", Method: http.MethodGet, Handler: handlers.ArtistAlbums},
}
