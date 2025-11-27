package main

import (
	"fmt"
	"net/http"

	"github.com/reecewilliams7/go-security-tools/cmd/gst-api/models"

	jwks "github.com/reecewilliams7/go-security-tools/internal/jsonWebKeys"
	ccs "github.com/reecewilliams7/go-security-tools/pkg/clientCredentials"
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

	clientIdCreator := ccs.NewUUIDv7ClientIdCreator()
	clientSecretCreator := ccs.NewCryptoRandClientSecretCreator()

	ci, err := clientIdCreator.Create()
	if err != nil {
		app.serverError(w, err)
	}

	cs, err := clientSecretCreator.Create()
	if err != nil {
		app.serverError(w, err)
	}

	co := models.ClientCredentialsOutput{ClientId: ci, ClientSecret: cs}

	app.json(w, co)
}
