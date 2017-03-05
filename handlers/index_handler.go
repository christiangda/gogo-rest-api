package handlers

import "net/http"

func Index(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Index"))
}
