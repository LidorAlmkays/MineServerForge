INSERT INTO minecraft_server (
    owner_email, server_name, allocated_ram_mb, max_player_amount, modes_names, plugins_names
) VALUES (
    :owner_email, :server_name, :allocated_ram_mb, :max_player_amount, :modes_names, :plugins_names
) RETURNING id;