-- name: CreateServerConfig :one
INSERT INTO minecraft_server_configs (owner_email, server_name, allocated_ram_mb, max_player_amount, modes_names, plugins_names)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;