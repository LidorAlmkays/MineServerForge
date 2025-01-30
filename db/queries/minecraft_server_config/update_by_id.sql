
-- name: UpdateServerConfig :exec
UPDATE minecraft_server_configs
SET owner_email = $1, server_name = $2, allocated_ram_mb = $3, max_player_amount = $4, modes_names = $5, plugins_names = $6
WHERE id = $7;

