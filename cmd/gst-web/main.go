package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

//go:embed templates/*.html
var templateFS embed.FS

type server struct {
	templates map[string]*template.Template
}

func newServer() *server {
	funcs := template.FuncMap{
		"add": func(a, b int) int { return a + b },
	}

	parse := func(pages ...string) *template.Template {
		files := append([]string{"templates/base.html"}, pages...)
		return template.Must(template.New("").Funcs(funcs).ParseFS(templateFS, files...))
	}

	return &server{
		templates: map[string]*template.Template{
			"index": parse("templates/index.html"),
			"jwk":   parse("templates/jwk.html"),
			"cc":    parse("templates/cc.html"),
		},
	}
}

func main() {
	srv := newServer()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", srv.handleIndex)
	mux.HandleFunc("GET /jwk", srv.handleJWKGet)
	mux.HandleFunc("POST /jwk", srv.handleJWKPost)
	mux.HandleFunc("GET /client-credentials", srv.handleCCGet)
	mux.HandleFunc("POST /client-credentials", srv.handleCCPost)

	addr := ":8080"
	log.Printf("Starting GST Web UI on http://localhost%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
