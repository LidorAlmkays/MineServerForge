package queries

import (
	"embed"
	"errors"
)

// Embed all SQL files inside the `minecraft_server_config` directories.
//
//go:embed minecraft_server_config/*.sql postgres_basic_queries/*.sql
var FS embed.FS

// GetQuery retrieves the contents of a query file from the embedded filesystem.
func GetQuery(path string) (string, error) {
	data, err := FS.ReadFile(path)
	if err != nil {
		return "", errors.New("query not found: " + path)
	}
	return string(data), nil
}
