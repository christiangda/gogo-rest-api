package controllers

import "net/http"

// Middleware
func Login(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Test"))
}

func Index(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Index"))
}