package handlers

import "net/http"

func Artist(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Test"))
}

func Artists(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Test"))
}

func ArtistAlbums(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Test"))
}
