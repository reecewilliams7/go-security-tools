package main

import (
	"fmt"
	"net/http"

	ccs "github.com/reecewilliams7/go-security-tools/internal/clientCredentials"
	jwks "github.com/reecewilliams7/go-security-tools/internal/jsonWebKeys"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	fmt.Fprint(w, "healthy")
}

func (app *application) jwk(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		app.methodNotAllowed(w)
		return
	}

	jwkc := jwks.NewRsaJsonWebKeyCreator()
	o, err := jwkc.Create()
	if err != nil {
		app.serverError(w, err)
	}

	app.json(w, o)
}

func (app *application) clientCredentials(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		app.methodNotAllowed(w)
		return
	}

	clientIdCreator := ccs.NewGuidClientIdCreator()
	clientSecretCreator := ccs.NewCryptoRandClientSecretCreator()
	cc := ccs.NewClientCredentialsCreator(clientIdCreator, clientSecretCreator)
	o, err := cc.Create()
	if err != nil {
		app.serverError(w, err)
	}

	app.json(w, o)
}
