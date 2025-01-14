package models

type MinecraftServerConfigModel struct {
	ID              int      `json:"id"`                // Unique primary key
	OwnerEmail      string   `json:"owner_email"`       // Email of the server owner
	ServerName      string   `json:"server_name"`       // Name of the server
	AllocatedRamMB  int      `json:"allocated_ram_mb"`  // RAM allocated to the server in MB
	MaxPlayerAmount int      `json:"max_player_amount"` // Maximum number of players
	ModesNames      []string `json:"modes_names"`       // List of mode names
	PluginsNames    []string `json:"plugins_names"`     // List of plugin names
}

func NewMinecraftServerConfigModel(OwnerEmail, ServerName string, AllocatedRamMB, MaxPlayerAmount int, ModesNames, PluginsNames []string) *MinecraftServerConfigModel {
	return &MinecraftServerConfigModel{OwnerEmail: OwnerEmail, ServerName: ServerName, AllocatedRamMB: AllocatedRamMB,
		MaxPlayerAmount: MaxPlayerAmount, ModesNames: ModesNames, PluginsNames: PluginsNames}
}
