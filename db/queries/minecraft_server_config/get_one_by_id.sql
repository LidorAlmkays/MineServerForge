-- name: GetServerConfigByID :one
SELECT id, owner_email, server_name, allocated_ram_mb, max_player_amount, modes_names, plugins_names
FROM minecraft_server_configs WHERE id = $1;
