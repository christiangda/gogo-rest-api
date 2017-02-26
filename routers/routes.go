package routers

import (
	"net/http"

	"github.com/christiangda/gogo-rest-api/controllers"
)

// Public:  Routes that doesn't need Authorization
var publicRoutes = Routes{
	{Name: "Index", Pattern: "/", Method: http.MethodGet, Handler: controllers.Index},
	{Name: "Login", Pattern: "/token-auth", Method: http.MethodPost, Handler: controllers.Login},
}

// Privated: Albums Routes
var albumsRoutes = Routes{
	{Name: "Album", Pattern: "/v1/albums/{id}", Method: http.MethodGet, Handler: controllers.Index},
	{Name: "Albums", Pattern: "/v1/albums?ids={ids}", Method: http.MethodGet, Handler: controllers.Login},
	{Name: "AlbumTracks", Pattern: "/v1/albums/{id}/tracks", Method: http.MethodGet, Handler: controllers.Login},
}

// Privated: Artists Routes
var artistsRoutes = Routes{
	{Name: "Artist", Pattern: "/v1/artists/{id}", Method: http.MethodGet, Handler: controllers.Index},
	{Name: "Artists", Pattern: "/v1/artists?ids={ids}", Method: http.MethodGet, Handler: controllers.Login},
	{Name: "ArtistAlbums", Pattern: "/v1/artists/{id}/albums", Method: http.MethodGet, Handler: controllers.Login},
}
