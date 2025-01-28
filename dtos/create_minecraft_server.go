package dtos

type CreateMinecraftServerDTO struct {
	OwnerEmail      string   `json:"ownerEmail"`
	ServerName      string   `json:"serverName"`
	MaxPlayerAmount int      `json:"maxPlayerAmount"`
	ModesNames      []string `json:"modesNames"`
	PluginsNames    []string `json:"pluginsNames"`
	AllocatedRamMB  int      `json:"allocatedRamMB"`
}

type CreateMinecraftServerResponseDTO struct {
	Id int64 `json:"id"`
}
