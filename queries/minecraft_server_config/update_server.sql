UPDATE minecraft_server
SET 
    owner_email = :owner_email,
    server_name = :server_name,
    allocated_ram_mb = :allocated_ram_mb,
    max_player_amount = :max_player_amount,
    modes_names = :modes_names,
    plugins_names = :plugins_names
WHERE id = :id;
