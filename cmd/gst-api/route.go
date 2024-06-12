package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/jwk", app.jwk)
	mux.HandleFunc("/client-credentials", app.clientCredentials)

	return mux
}
