CREATE TABLE minecraft_server (
    id SERIAL PRIMARY KEY,                    
    owner_email VARCHAR(255) NOT NULL,        
    server_name VARCHAR(255) NOT NULL,  
    max_player_amount INT NOT NULL,           
    modes_names TEXT[] NOT NULL,              
    plugins_names TEXT[] NOT NULL,  
    allocated_ram_mb INT NOT NULL,          
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  
);