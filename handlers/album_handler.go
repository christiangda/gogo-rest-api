package handlers

import "net/http"

func Album(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Test"))
}
