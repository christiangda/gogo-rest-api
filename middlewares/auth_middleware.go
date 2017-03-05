package middlewares

import "net/http"

func authMiddleware(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	next(w, req)
}
