-- name: DeleteServerConfig :exec
DELETE FROM minecraft_server_configs WHERE id = $1;