package api

import (
	_ "embed"
	"net/http"
)

// Embed the Swagger files
//
//go:embed docs/swagger.yaml
var SwaggerYAML []byte

//go:embed docs/swagger.json
var SwaggerJSON []byte

// ServeSwagger serves the Swagger files over HTTP
func ServeSwagger(m *http.ServeMux) {
	m.HandleFunc("/swagger.yaml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/yaml")
		w.Write(SwaggerYAML)
	})

	m.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(SwaggerJSON)
	})

}
