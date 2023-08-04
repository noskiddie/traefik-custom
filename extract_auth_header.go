// extract_auth_header.go

package main

import (
	"net/http"
	"github.com/traefik/traefik/v2/pkg/log"
)

func ExtractAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		authHeader := req.Header.Get("Authorization")
		log.FromContext(req.Context()).Infof("Authorization header: %s", authHeader)

		// Perform your desired actions with the Authorization header (e.g., log stats, store data, etc.)

		next.ServeHTTP(rw, req)
	})
}
