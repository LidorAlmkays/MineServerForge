-- name: EnsureTableExists :exec
CREATE TABLE IF NOT EXISTS minecraft_server_configs (
    id SERIAL PRIMARY KEY,
    owner_email VARCHAR(100) NOT NULL,
    server_name VARCHAR(100) NOT NULL,
    allocated_ram_mb INT NOT NULL,
    max_player_amount INT NOT NULL,
    modes_names TEXT[],   -- Array of strings (PostgreSQL array type)
    plugins_names TEXT[]  -- Array of strings (PostgreSQL array type)
);
