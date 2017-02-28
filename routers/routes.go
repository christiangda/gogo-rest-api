package routers

import (
	"github.com/christiangda/gogo-rest-api/handlers"
	"net/http"
)

// Public:  Routes that doesn't need Authorization
var publicRoutes = Routes{
	{Name: "Index", Pattern: "/", Method: http.MethodGet, Handler: handlers.Index},
	{Name: "Login", Pattern: "/token-auth", Method: http.MethodPost, Handler: handlers.Login},
}

// Privated: Albums Routes
var albumsRoutes = Routes{
	{Name: "Album", Pattern: "/albums/{id}", Method: http.MethodGet, Handler: handlers.Album},
	{Name: "Albums", Pattern: "/albums?ids={ids}", Method: http.MethodGet, Handler: handlers.Login},
	{Name: "AlbumTracks", Pattern: "/albums/{id}/tracks", Method: http.MethodGet, Handler: handlers.Login},
}

// Privated: Artists Routes
var artistsRoutes = Routes{
	{Name: "Artist", Pattern: "/artists/{id}", Method: http.MethodGet, Handler: handlers.Index},
	{Name: "Artists", Pattern: "/artists?ids={ids}", Method: http.MethodGet, Handler: handlers.Login},
	{Name: "ArtistAlbums", Pattern: "/artists/{id}/albums", Method: http.MethodGet, Handler: handlers.Login},
}
